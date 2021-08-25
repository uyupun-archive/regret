package admin

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/database/query"
	"github.com/uyupun/regret/models"
)

func GetInquiryValidation(c echo.Context) error {
	serviceId := c.QueryParam("service_id")
	tmpServiceId, err := strconv.Atoi(serviceId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	inquiryValidation, err := query.GetInquiryValidationByServiceId(tmpServiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, inquiryValidation)
}

func EditInquiryValidation(c echo.Context) error {
	inquiryValidation := new(models.InquiryValidation)
	err := c.Bind(&inquiryValidation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = query.EditInquiryValidation(*inquiryValidation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}
