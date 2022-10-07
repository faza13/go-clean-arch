package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type User struct {
}

type IUser interface {
	Handle(next echo.HandlerFunc) echo.HandlerFunc
}

func NewUser() IUser {
	return &User{}
}

func (u *User) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonClaim := (c.Get("user").(*jwt.Token)).Claims.(jwt.MapClaims)
		userData := jsonClaim["data"]
		c.Set("user", userData)
		return next(c)
	}
}
