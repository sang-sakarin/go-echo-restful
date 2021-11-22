package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (ce *CustomEcho) UseSetting(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Static("/static", "static")

	ce.UseRoute(e)
}
