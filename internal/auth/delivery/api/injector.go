//go:build wireinject
// +build wireinject

package api

import (
	"base/app/config"
	"base/app/providers/database"
	"base/app/providers/hash"
	respositories2 "base/internal/auth/repositories"
	services2 "base/internal/auth/services"
	"base/internal/user/respositories"
	"base/utils/token"
	"base/utils/uuid"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeAuthApi(echo *echo.Echo, provider *database.Provider, JwtConfig *config.JWT) IAuthApi {
	wire.Build(
		uuid.NewUuidProvider,
		token.NewTokenProvider,
		hash.NewHashProvider,
		services2.NewAuthService,
		respositories.NewUserRepository,
		respositories2.NewRefreshTokenRepositories,
		NewAuthAPi,
	)
	return nil
}
