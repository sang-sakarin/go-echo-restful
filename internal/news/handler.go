package news

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo-restful/pkg/database"
	"go-echo-restful/pkg/response"
	"net/http"
)

// Create godoc
// @Summary
// @Description List news
// @Tags News
// @Router /v1/news [post]
// @Param title body SerializerForm true "abc"
// @Success 200 {object} SerializerForm "Test"
func Create(c echo.Context) error {
	n := new(SerializerForm)

	if err := c.Bind(n); err != nil {
		return err
	}

	if err := c.Validate(n); err != nil {
		return err
	}

	nn := News{Title: n.Title, Description: n.Description, IsPublish: n.IsPublish}

	result := database.DBConn.Create(&nn)

	fmt.Println(result.Error)

	return c.JSON(http.StatusCreated, response.Response{Success: true, Data: "WOW"})
}
