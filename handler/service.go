package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/models"
)

func AddService(c echo.Context) error {
	addService := new(models.AddService)
	err := c.Bind(&addService)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("%s\n", addService.Name)
	// gormでDBに追加
	return c.JSON(http.StatusOK, "test")
}
