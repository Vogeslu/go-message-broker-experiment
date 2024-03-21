package core

import (
	"message_broker/internal/config"
	"message_broker/internal/listener"
	"message_broker/internal/logger"
)

func StartBroker() {
	logger.Logger.Info().Msg("Starting message broker")

	config.LoadConfig()
	listener.StartListener()
}
