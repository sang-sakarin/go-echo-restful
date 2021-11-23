package pkg

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stoewer/go-strcase"
	"go-echo-restful/pkg/response"
	"net/http"
	"reflect"
	"strings"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme
	cv.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		n := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if n != "-" {
			return n
		}

		n = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if n != "-" {
			return n
		}
		return ""
	})

	if err := cv.validator.Struct(i); err != nil {

		errs := ""
		for _, err := range err.(validator.ValidationErrors) {
			errs += fmt.Sprintf("%s: %s\r\n", strcase.SnakeCase(err.Field()), err.ActualTag())
		}

		res := response.Response{Success: false, Message: errs, Data: nil, Total: 0}

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	return nil
}
