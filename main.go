package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-restful/controllers"
	"net/http"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "It's me!!!")
	})

	v1 := e.Group("v1")
	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.GetUser)
	v1.POST("/users", controllers.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
