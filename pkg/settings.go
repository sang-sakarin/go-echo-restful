package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-restful/pkg/middleware"
)

func (ce *CustomEcho) UseSetting(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Abc)

	e.Static("/static", "static")

	ce.UseRoute(e)
}
