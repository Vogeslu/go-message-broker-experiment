package core

import (
	"message_broker/internal/command"
	"message_broker/internal/config"
	"message_broker/internal/listener"
	"message_broker/internal/logger"
)

func StartBroker() {
	logger.Logger.Info().Msg("Starting message broker")

	config.LoadConfig()
	command.RegisterCommands()
	listener.StartListener()
}
