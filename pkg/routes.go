package pkg

import (
	"github.com/labstack/echo/v4"
	"go-echo-restful/internal/user"
	"net/http"
)

func (ce *CustomEcho) UseRoute(e *echo.Echo) {

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "It's me!!!")
	})

	v1 := e.Group("v1")
	v1.GET("/users", user.List)
	v1.GET("/users/:id", user.Retrieve)
	v1.POST("/users", user.Create)

}
