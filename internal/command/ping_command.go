package command

import "message_broker/internal/session"

type PingCommand struct{}

func (ping *PingCommand) Name() string {
	return "ping"
}

func (ping *PingCommand) Description() string {
	return "Send ping message"
}

func (ping *PingCommand) Arguments() []Argument {
	return nil
}

func (ping *PingCommand) CanBeCalled(session *session.Session) error {
	return nil
}

func (ping *PingCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (ping *PingCommand) ParsePayload(*session.Session, map[string]*string) (error, interface{}) {
	return nil, nil
}

func (ping *PingCommand) OnCalled(*session.Session, interface{}) (error, interface{}) {
	return nil, nil
}

func (ping *PingCommand) GetResponse(*session.Session, interface{}) []string {
	return []string{"pong"}
}
