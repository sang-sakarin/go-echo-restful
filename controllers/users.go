package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	User struct {
		Name         string  `json:"name" form:"name"`
		Description  *string `json:"descriptions" form:"description1" validate:"required""`
		Description2 *string `json:"new" form:"description2" validate:"required"`
	}
)

func GetUsers(c echo.Context) error {
	team := c.QueryParam("team")

	return c.String(http.StatusOK, "team: "+team)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func CreateUser(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
