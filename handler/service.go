package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"github.com/uyupun/regret/database/query"
	"github.com/uyupun/regret/models"
)

func AddService(c echo.Context) error {
	service := new(models.Service)
	err := c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	service.AccessToken = randstr.String(20)

	err = query.AddService(*service)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "test")
}
