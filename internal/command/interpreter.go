package command

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"message_broker/internal/session"
)

var currentCommand *Command = nil
var arguments map[string]*string = nil

func ParseLineInput(session *session.Session, line string) (error, []Argument, []string) {
	newCommand := currentCommand == nil

	if currentCommand == nil {
		err := handleNewCommand(session, line)
		if err != nil {
			return err, nil, nil
		}
	}

	command := *currentCommand

	missingArguments := areArgumentsProvided(currentCommand, arguments)

	if newCommand && len(missingArguments) > 0 {
		return nil, missingArguments, nil
	}

	if len(missingArguments) > 0 {
		currentArgument := missingArguments[0]

		err := handleArgument(session, line, command, currentArgument)
		if err != nil {
			return err, missingArguments, nil
		}

		missingArguments = areArgumentsProvided(currentCommand, arguments)
	}

	if len(missingArguments) == 0 {
		err, response := processCommand(session, command)
		if err != nil {
			return err, nil, nil
		}

		return nil, nil, response
	}

	return nil, missingArguments, nil
}

func handleNewCommand(session *session.Session, line string) error {
	err, command := findCommand(line)
	if err != nil {
		return err
	}

	err = (*command).CanBeCalled(session)
	if err != nil {
		return err
	}

	currentCommand = command
	arguments = nil

	return nil
}

func handleArgument(session *session.Session, line string, command Command, currentArgument Argument) error {
	var input *string = nil
	logger.Logger.Debug().Msg(line)
	if len(line) > 0 {
		input = &line
	} else if currentArgument.Required {
		return errors.New(fmt.Sprintf("Value for %s is required", currentArgument.Name))
	}

	err := command.ValidateArgument(session, currentArgument, input)
	if err != nil {
		return err
	}

	if arguments == nil {
		arguments = map[string]*string{}
	}

	arguments[currentArgument.Key] = input

	return nil
}

func processCommand(session *session.Session, command Command) (error, []string) {
	defer cleanCurrentCommand()

	err, payload := command.ParsePayload(session, arguments)
	if err != nil {
		return err, nil
	}

	err, result := command.OnCalled(session, payload)
	if err != nil {
		return err, nil
	}

	response := command.GetResponse(session, result)

	return nil, response
}

func cleanCurrentCommand() {
	currentCommand = nil
	arguments = nil
}
