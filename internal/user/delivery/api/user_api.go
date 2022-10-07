package api

import (
	"base/app/common"
	"base/app/constants"
	"base/internal/user/resources"
	"base/internal/user/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserApi struct {
	userService services.IUserService
}

type IUserApi interface {
	Store(echo.Context) error
}

func NewUserApi(echo *echo.Group, service services.IUserService) IUserApi {
	userApi := &UserApi{
		userService: service,
	}
	g := echo.Group("/v1")
	g.POST("/user", userApi.Store)
	g.POST("/user/change-password", userApi.ChangePassword)

	return userApi
}

func (u UserApi) Store(c echo.Context) error {
	userStoreRequest := &resources.UserStoreRequest{}
	response := common.Response{IsSuccess: false}

	err := c.Bind(&userStoreRequest)
	if err != nil {
		response.Message = "Terjadi kesalahan pada server coba beberapa saat lagi"
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := u.userService.Store(userStoreRequest)

	if err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = res
	return c.JSON(http.StatusOK, response)
}

func (u UserApi) ChangePassword(c echo.Context) error {
	var response common.Response
	changePassReq := new(resources.ChangePasswordRequest)
	c.Bind(changePassReq)
	errs := constants.Validate.Make(changePassReq)
	if errs != nil {
		return response.Error422(c, "data tidak valid", errs)
	}
	user := c.Get("user").(map[string]interface{})
	userID := uint64(user["id"].(float64))

	err := u.userService.ChangePassword(userID, changePassReq.NewPassword)
	if err != nil {
		return err
	}

	changePassReq = nil
	return response.Success(c, map[string]string{}, "Ganti password berhasi")
}
