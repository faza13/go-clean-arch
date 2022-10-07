//go:build wireinject
// +build wireinject

package api

import (
	"base/app/providers/database"
	"base/app/providers/hash"
	"base/internal/user/respositories"
	"base/internal/user/services"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeUserApi(echo *echo.Group, provider *database.Provider) IUserApi {
	wire.Build(hash.NewHashProvider, services.NewUserService, respositories.NewUserRepository, NewUserApi)
	return nil
}
