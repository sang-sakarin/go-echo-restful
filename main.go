package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "It's me!!!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
