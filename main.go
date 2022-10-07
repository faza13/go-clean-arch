package main

import (
	"base/app/config"
	"base/app/constants"
	"base/app/providers/Validation"
	"base/app/providers/cache"
	"base/app/providers/database"
	"base/app/providers/translator"
	_authApi "base/internal/auth/delivery/api"
	_userApi "base/internal/user/delivery/api"
	middleware2 "base/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {
	appConfig := config.NewApp()

	loc, _ := time.LoadLocation(appConfig.TimeZone)
	// handle err
	time.Local = loc

	constants.Trans = translator.NewTranslatorProvider("id")
	constants.Validate = Validation.NewValidatroProvider(constants.Trans)

	e := echo.New()
	e.Use(middleware.CORS())

	dbProvider := database.NewMariaDBProvider(appConfig.Database)
	cache.NewRedisProvider(appConfig.Cache)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ape lo liat2")
	})

	_authApi.InitializeAuthApi(e, dbProvider, appConfig.JWT)

	userMiddleware := middleware2.NewUser()
	auth := e.Group("", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(appConfig.JWT.JwtSecret),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		//ContextKey:  "data",
	}), userMiddleware.Handle)

	_userApi.InitializeUserApi(auth, dbProvider)

	e.POST("healt-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "ape lo liat2")
	})
	e.Logger.Fatal(e.Start(":" + appConfig.Port))
}
