package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-restful/pkg"
)

func main() {
	e := echo.New()

	ce := pkg.CustomEcho{}

	ce.UseSetting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
