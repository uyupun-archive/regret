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

	err := slack.PostWebhook(url, &slack.WebhookMessage{
		Text: ":rabbit:問い合わせを受信したぺこ:rabbit2:",
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewDividerBlock(),
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: ":rabbit:問い合わせが届いたぺこ！:rabbit2:",
					},
				},
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: fmt.Sprintf("> タイトル: %s", inquiry.Title),
					},
				},
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: fmt.Sprintf("> メールアドレス: %s", inquiry.Email),
					},
				},
				&slack.SectionBlock{
					Type: slack.MBTSection,
					Text: &slack.TextBlockObject{
						Type: "mrkdwn",
						Text: fmt.Sprintf("> 本文:\n%s", inquiry.Message),
					},
				},
				slack.NewDividerBlock(),
			},
		},
	})
	return err
}
