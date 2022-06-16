package client

import (
	"encoding/json"
	"go-slack-message-client/internal/logger"
	"go-slack-message-client/internal/models"

	"github.com/gin-gonic/gin"
)

var Message models.Messages
var PayloadText string
var PayloadTS string

func GetPayloadFrontEnd(c *gin.Context) {

	logger, dispose := logger.New()
	defer dispose()

	body := models.Messages{}
	decoder := json.NewDecoder(c.Request.Body)
	logger.Info("Getting Payload" + body.PayloadTS)
	if err := decoder.Decode(&body); err != nil {
		logger.Error(err.Error())
		return
	}
	PayloadText = body.PayloadText
	PayloadTS = body.PayloadTS
	logger.Info("PayloadText: " + PayloadText + " PayloadTS: " + PayloadTS)

	ReplyMessage(PayloadTS, PayloadText)
}
