package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/uyupun/regret/database/query"
)

func AccessTokenVerification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"][0]
		clientAccessToken := getClientAccessToken(authorization)

		isCorrect, err := verifyAccessToken(clientAccessToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if !isCorrect {
			return c.JSON(http.StatusBadRequest, "{}")
		}

		return next(c)
	}
}

func getClientAccessToken(authorization string) string {
	ary := strings.Split(authorization, " ")
	appKey := ary[1]
	return appKey
}

func verifyAccessToken(clientAccessToken string) (bool, error) {
	serverAccessToken, err := getServerAccessToken(clientAccessToken)
	if err != nil {
		return false, err
	}

	if clientAccessToken == serverAccessToken {
		return true, nil
	}
	return false, nil
}

func getServerAccessToken(clientAccessToken string) (string, error) {
	service, err := query.GetServiceByAccessToken(clientAccessToken)
	if err != nil {
		return "", err
	}
	serverAccessToken := service.AccessToken
	return serverAccessToken, nil
}
