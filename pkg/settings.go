package pkg

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-restful/pkg/config"
	"go-echo-restful/pkg/database"
	"go-echo-restful/pkg/middleware"
)

func (ce *CustomEcho) UseSetting(e *echo.Echo) {

	env.Parse(&config.Cfg)

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Abc)

	e.Static("/static", "static")

	ce.UseRoute(e)

	database.UseDatabase()
}
