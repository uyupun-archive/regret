package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddService(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")
}
