package command

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"message_broker/internal/session"
)

type GreetCommand struct{}

func (greet *GreetCommand) Name() string {
	return "greet"
}

func (greet *GreetCommand) Description() string {
	return "Greet to server"
}

func (greet *GreetCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "name",
			Name:     "Name",
			Help:     "Name of Client",
			Required: true,
		},
	}
}

func (greet *GreetCommand) CanBeCalled(session *session.Session) error {
	return nil
}

func (greet *GreetCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (greet *GreetCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	return nil, data["name"]
}

func (greet *GreetCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	if session.Greeted {
		return errors.New("Connection already greeted"), nil
	}

	name := payload.(*string)

	session.Name = name
	session.Greeted = true

	logger.Logger.Info().Msg(fmt.Sprintf("Session %s greeted with name %s", session.Conn.RemoteAddr().String(), *session.Name))

	return nil, name
}

func (greet *GreetCommand) GetResponse(session *session.Session, payload interface{}) []string {
	name := payload.(*string)

	return []string{fmt.Sprintf("Greeted as %s", *name)}
}
