package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/database/query"
)

func GetCategories(c echo.Context) error {
	serviceId := c.QueryParam("service_id")
	tmpServiceId, err := strconv.Atoi(serviceId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	categories, err := query.GetCategories(tmpServiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, categories)
}
