package api

import (
	"base/app/common"
	"base/internal/auth/resources"
	"base/internal/auth/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthApi struct {
	authService services.IAuthService
}

type IAuthApi interface {
	Store(echo.Context) error
}

func NewAuthAPi(echo *echo.Echo, authService services.IAuthService) IAuthApi {
	authApi := &AuthApi{
		authService: authService,
	}
	g := echo.Group("/v1")
	g.POST("/login/:entity", authApi.Store)
	g.POST("/validate", authApi.Validate)
	g.POST("/refresh", authApi.RefreshToken)

	return authApi
}

func (u AuthApi) Store(c echo.Context) error {
	loginRequest := &resources.LoginRequest{}
	response := common.Response{IsSuccess: false}

	platform := c.Request().Header.Get("X-Platform")

	err := c.Bind(&loginRequest)
	if err != nil {
		response.Message = "Terjadi kesalahan pada server coba beberapa saat lagi"
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := u.authService.Store(loginRequest.Username, loginRequest.Password, 0, platform)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Data = token

	return c.JSON(http.StatusOK, response)
}

func (u AuthApi) RefreshToken(c echo.Context) error {
	response := common.Response{IsSuccess: false}
	var refreshRequest resources.RefreshRequest

	err := c.Bind(&refreshRequest)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data, err = u.authService.GenerateNewToken(refreshRequest.RefreshToken)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusBadRequest, response)
}

func (u AuthApi) Validate(c echo.Context) error {
	response := common.Response{IsSuccess: false}
	token := c.FormValue("token")
	data, err := u.authService.Validate(token)

	if err != nil {
		response.Message = err.Error()

		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = map[string]interface{}{"token": data}

	return c.JSON(http.StatusOK, response)
}
