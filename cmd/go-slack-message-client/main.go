package main

import (
	"go-slack-message-client/internal/logger"
	"go-slack-message-client/internal/routes"
	"go-slack-message-client/internal/slacktools"
)

func main() {

	logger, dispose := logger.New()
	defer dispose()

	logger.Info("Starting Worker")
	go slacktools.SlackSocket()

	logger.Info("Starting API")
	routes.MakeDefaultRoutes()

}
