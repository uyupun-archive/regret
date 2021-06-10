package general

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
	"github.com/uyupun/regret/database/query"
	"github.com/uyupun/regret/models"
)

func PostInquiry(c echo.Context) error {
	inquiry := new(models.Inquiry)
	err := c.Bind(&inquiry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = postSlack(*inquiry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "{}")
}

func postSlack(inquiry models.Inquiry) error {
	category, err := query.GetCategoryById(inquiry.CategoryId)
	if err != nil {
		return err
	}

	url := os.Getenv("SLACK_INCOMING_WEBHOOK_URL")

	notifyBlock := makeSlackSectionBlock(":rabbit:問い合わせが届いたぺこ！:rabbit2:")
	subjectBlock := makeSlackSectionBlock(fmt.Sprintf("> 件名: %s", inquiry.Subject))
	emailBlock := makeSlackSectionBlock(fmt.Sprintf("> メールアドレス: %s", inquiry.Email))
	categoryBlock := makeSlackSectionBlock(fmt.Sprintf("> カテゴリ: %s(%s)", category.Name, category.NameJa))
	textBlock := makeSlackSectionBlock(fmt.Sprintf("> 本文:\n> %s", inquiry.Text))

	err = slack.PostWebhook(url, &slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewDividerBlock(),
				notifyBlock,
				subjectBlock,
				emailBlock,
				categoryBlock,
				textBlock,
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
