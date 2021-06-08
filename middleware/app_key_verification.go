package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func AppKeyVerification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"][0]
		clientAppKey := getClientAppKey(authorization)

		isCorrect, err := verifyAppKey(clientAppKey)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if !isCorrect {
			return c.JSON(http.StatusBadRequest, "{}")
		}
		return next(c)
	}
}

func getClientAppKey(authorization string) string {
	ary := strings.Split(authorization, " ")
	appKey := ary[1]
	return appKey
}

func verifyAppKey(clientAppKey string) (bool, error) {
	serverAppKey := os.Getenv("APP_KEY")

	if clientAppKey == serverAppKey {
		return true, nil
	}
	return false, nil
}
