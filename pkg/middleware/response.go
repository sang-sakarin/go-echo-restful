package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Abc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("begin")
		a := next(c)
		fmt.Println("end")

		return a
	}
}
