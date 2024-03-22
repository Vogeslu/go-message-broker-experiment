package command

import (
	"fmt"
	"message_broker/internal/session"
)

type HelpCommand struct{}

func (help *HelpCommand) Name() string {
	return "help"
}

func (help *HelpCommand) Description() string {
	return "List available commands"
}

func (help *HelpCommand) Arguments() []Argument {
	return nil
}

func (help *HelpCommand) CanBeCalled(session *session.Session) error {
	return nil
}

func (help *HelpCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (help *HelpCommand) ParsePayload(*session.Session, map[string]*string) (error, interface{}) {
	return nil, nil
}

func (help *HelpCommand) OnCalled(*session.Session, interface{}) (error, interface{}) {
	return nil, nil
}

func (help *HelpCommand) GetResponse(*session.Session, interface{}) []string {
	var output []string

	for i := range commands {
		command := commands[i]
		output = append(output, fmt.Sprintf("%s (%s)", command.Name(), command.Description()))
	}

	return output
}
