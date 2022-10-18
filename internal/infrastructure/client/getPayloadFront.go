package client

import (
	"encoding/json"
	slackmessagesapi "slack-message-api/internal/domain/slackmessagesapi"
	"slack-message-api/internal/infrastructure/simlogger"

	"github.com/gin-gonic/gin"
)

var Message slackmessagesapi.Messages
var PayloadText string
var PayloadTS string
var Replied bool
var ID int

func GetPayloadFrontEnd(c *gin.Context) {

	logger, dispose := simlogger.New()
	defer dispose()

	body := slackmessagesapi.Messages{}
	decoder := json.NewDecoder(c.Request.Body)
	logger.Info("Getting Payload" + body.PayloadTS)
	if err := decoder.Decode(&body); err != nil {
		logger.Error(err.Error())
		return
	}
	PayloadText = body.PayloadText
	PayloadTS = body.PayloadTS
	Replied = body.Replied
	ID = body.ID

	logger.Info("PayloadText: " + PayloadText + " PayloadTS: " + PayloadTS)

	ReplyMessage(PayloadTS, PayloadText, Replied, ID)
}
