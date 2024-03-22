package command

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"strings"
)

var commands []Command

func RegisterCommands() {
	commands = []Command{
		&PingCommand{},
		&HelpCommand{},

		&GreetCommand{},

		&ListSessionsCommand{},

		&SubscribeTopicCommand{},
		&UnsubscribeTopicCommand{},
		&ListSubscriptionsCommand{},
		&ListSubscribersCommand{},
		&ListTopicsCommand{},

		&SendMessageCommand{},
		&AcknowledgeMessageCommand{},
		&ListAcknowledgementsCommand{},
	}

	var names []string

	for i := range commands {
		names = append(names, commands[i].Name())
	}

	logger.Logger.Info().Msg(fmt.Sprintf("Registered commands %s", strings.Join(names, ", ")))
}

func findCommand(name string) (error, *Command) {
	for i := range commands {
		command := commands[i]
		if command.Name() == name {
			return nil, &command
		}
	}

	return errors.New("Command not found"), nil
}
