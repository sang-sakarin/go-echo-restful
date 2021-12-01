package user

import (
	"github.com/labstack/echo/v4"
	"go-echo-restful/pkg/config"
	"go-echo-restful/pkg/response"
	"net/http"
)

type (
	User struct {
		Name         string  `json:"name" form:"name"`
		Description  *string `json:"descriptions" form:"description1" validate:"required"`
		Description2 *string `json:"new" form:"description3" validate:"required"`
		LongName     string  `json:"long_name" validate:"required"`
	}
)

func List(c echo.Context) error {
	return c.JSON(http.StatusOK, response.Response{Success: true, Data: config.Cfg})
}

func Retrieve(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, response.Response{Success: true, Data: id})
}

func Create(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response.Response{Success: true, Data: u})
}
