package common

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
	Trans    ut.Translator
)

type Response struct {
	Message   string      `json:"message"`
	Errors    interface{} `json:"errors"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"success"`
}

func (g *Response) Error422(context echo.Context, message string, Errors interface{}) error {
	g.Message = message
	g.Errors = Errors

	return context.JSON(http.StatusUnprocessableEntity, g)
}

func (g *Response) Error400(context echo.Context, message string) error {
	g.Message = message

	return context.JSON(http.StatusBadRequest, g)
}

func (g *Response) Success(context echo.Context, data interface{}, message string) error {
	g.Message = message
	if g.Message == "" {
		g.Message = "OK"
	}

	g.Data = data
	g.IsSuccess = true

	return context.JSON(http.StatusOK, g)
}
