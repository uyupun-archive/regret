package general

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	err = postSlack(*inquiry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func postSlack(inquiry models.Inquiry) error {
	url := os.Getenv("SLACK_INCOMING_WEBHOOK_URL")

	title := makeSlackAttachment("タイトル", inquiry.Title)
	email := makeSlackAttachment("メールアドレス", inquiry.Email)
	category := makeSlackAttachment("カテゴリ", strconv.Itoa(inquiry.CategoryId))
	message := makeSlackAttachment("本文", inquiry.Message)

	err := slack.PostWebhook(url, &slack.WebhookMessage{
		Username:  "test",
		IconEmoji: ":rabbit:",
		Text:      ":rabbit:問い合わせを受信したぺこ:rabbit2:",
		Attachments: []slack.Attachment{
			title,
			email,
			category,
			message,
		},
	})
	return err
}

func makeSlackAttachment(title string, value string) slack.Attachment {
	titleField := slack.AttachmentField{Title: title + ":", Value: value}
	attachment := slack.Attachment{}
	attachment.Fields = []slack.AttachmentField{titleField}
	color := "good"
	attachment.Color = color
	return attachment
}
