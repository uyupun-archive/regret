package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"github.com/uyupun/regret/database/query"
	"github.com/uyupun/regret/models"
)

func GetServices(c echo.Context) error {
	services, err := query.GetServices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, services)
}

func AddService(c echo.Context) error {
	service := new(models.Service)
	err := c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	service.AccessToken = randstr.String(20)

	addedService, err := query.AddService(*service)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = query.AddInquiryValidation(addedService.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func EditService(c echo.Context) error {
	service := new(models.Service)
	err := c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = query.EditService(*service)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func DeleteService(c echo.Context) error {
	service := new(models.Service)
	err := c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = query.DeleteService(*service)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}
