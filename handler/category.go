package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/database/query"
	"github.com/uyupun/regret/models"
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

func AddCategory(c echo.Context) error {
	category := new(models.Category)
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = query.AddCategory(*category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}
