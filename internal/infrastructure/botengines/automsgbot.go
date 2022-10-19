package slackclient

import (
	"slack-messages-api/internal/infrastructure/environment"
	"slack-messages-api/internal/infrastructure/logger"

	"github.com/slack-go/slack"
)

func BotReply(PayloadText string, PayloadTS string) {
	logger, dispose := logger.New()
	defer dispose()

	env := environment.GetInstance()
	token := env.SLACK_AUTH_TOKEN
	channel := env.SLACK_CHANNEL_ID
	api := slack.New(token)
	attachment := slack.Attachment{
		Color: "green",
		Text:  "Olá, eu sou o Re. Em breve ficarei mais inteligente. :smile:",
	}

	logger.Info("AUTO_REPLY - Sending Message")
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
	logger.Info("AUTO_REPLY - Message Successfully Sent to" + channelID + " at " + timestamp)
}
