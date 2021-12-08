package pkg

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	//echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-echo-restful/docs"
	"go-echo-restful/internal/news"
	"go-echo-restful/internal/user"
	"net/http"
)

func (ce *CustomEcho) UseRoute(e *echo.Echo) {

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "It's me!!!")
	})

	e.GET("/docs/*", echoSwagger.WrapHandler)

	v1 := e.Group("v1")
	v1.GET("/users", user.List)
	v1.GET("/users/:id", user.Retrieve)
	v1.POST("/users", user.Create)

	v1.POST("/news", news.Create)

}
