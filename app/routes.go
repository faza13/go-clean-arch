package app

import (
	"github.com/labstack/echo/v4"
)

type router struct {
	Echo *echo.Echo
}

type IRouter interface {
	GetRouter() *echo.Echo
}

func NewRouters(e *echo.Echo) IRouter {
	return router{
		Echo: e,
	}
}

func (r router) GetRouter() *echo.Echo {
	return r.Echo
}
