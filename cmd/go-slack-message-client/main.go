package main

import (
	"slack-message-api/internal/infrastructure/routes"
	"slack-message-api/internal/infrastructure/simlogger"
	"slack-message-api/internal/infrastructure/slacktools"
)

func main() {

	logger, dispose := simlogger.New()
	defer dispose()

	logger.Info("Starting Worker")
	go slacktools.SlackSocket()

	logger.Info("Starting API")
	routes.MakeDefaultRoutes()

}
