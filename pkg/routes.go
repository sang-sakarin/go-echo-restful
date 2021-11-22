package pkg

import (
	"github.com/labstack/echo/v4"
	"go-echo-restful/controllers"
	"net/http"
)

func (ce *CustomEcho) UseRoute(e *echo.Echo) {

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "It's me!!!")
	})

	v1 := e.Group("v1")
	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.GetUser)
	v1.POST("/users", controllers.CreateUser)

}
