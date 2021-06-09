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

	err = postSlack(*inquiry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func postSlack(inquiry models.Inquiry) error {
	url := os.Getenv("SLACK_INCOMING_WEBHOOK_URL")

	notify := makeSlackSectionBlock(":rabbit:問い合わせが届いたぺこ！:rabbit2:")
	subject := makeSlackSectionBlock(fmt.Sprintf("> 件名: %s", inquiry.Subject))
	email := makeSlackSectionBlock(fmt.Sprintf("> メールアドレス: %s", inquiry.Email))
	text := makeSlackSectionBlock(fmt.Sprintf("> 本文:\n> %s", inquiry.Text))

	err := slack.PostWebhook(url, &slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewDividerBlock(),
				notify,
				subject,
				email,
				text,
				slack.NewDividerBlock(),
			},
		},
	})
	return err
}

func makeSlackSectionBlock(text string) *slack.SectionBlock {
	return &slack.SectionBlock{
		Type: slack.MBTSection,
		Text: &slack.TextBlockObject{
			Type: "mrkdwn",
			Text: text,
		},
	}
}
