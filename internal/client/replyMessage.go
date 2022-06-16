package client

import (
	"go-slack-message-client/internal/environment"
	"go-slack-message-client/internal/logger"

	"github.com/slack-go/slack"
)

func ReplyMessage(PayloadTS string, PayloadText string) {

	logger, dispose := logger.New()
	defer dispose()

	env := environment.GetInstance()
	token := env.SLACK_AUTH_TOKEN
	channel := env.SLACK_CHANNEL_ID
	api := slack.New(token)
	attachment := slack.Attachment{
		Color: "green",
		Text:  PayloadText,
	}

	logger.Info("Sending Message")
	channelID, timestamp, err := api.PostMessage(
		channel,
		// slack.MsgOptionText(PayloadText, false),
		slack.MsgOptionTS(PayloadTS),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Message Successfully Sent to" + channelID + " at " + timestamp)
}
