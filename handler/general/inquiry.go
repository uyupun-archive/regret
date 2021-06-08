package general

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
	"github.com/uyupun/regret/models"
)

func PostInquiry(c echo.Context) error {
	inquiry := new(models.Inquiry)
	err := c.Bind(&inquiry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("%#v\n", inquiry)

	err = postSlack()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func postSlack() error {
	token := os.Getenv("SLACK_OAUTH_TOKEN")
	client := slack.New(token)
	text := slack.MsgOptionText("Hello World", true)
	channel := os.Getenv("SLACK_CHANNEL")
	_, _, err := client.PostMessage("#"+channel, text)
	return err
}
