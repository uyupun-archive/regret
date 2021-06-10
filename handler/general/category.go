package general

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/database/query"
)

func GetCategories(c echo.Context) error {
	authorization := c.Request().Header["Authorization"][0]
	clientAccessToken := getClientAccessToken(authorization)

	categories, err := query.GetCategoriesByAccessToken(clientAccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, categories)
}

func getClientAccessToken(authorization string) string {
	ary := strings.Split(authorization, " ")
	appKey := ary[1]
	return appKey
}
