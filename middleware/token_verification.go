package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func TokenVerification(next echo.HandlerFunc) echo.HandlerFunc {
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
	accessToken := ary[1]
	return accessToken
}

func verifyAccessToken(clientAccessToken string) (bool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return false, err
	}
	serverAccessToken := os.Getenv("ACCESS_TOKEN")

	if clientAccessToken == serverAccessToken {
		return true, nil
	}
	return false, nil
}
