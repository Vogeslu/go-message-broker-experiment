package command

import "message_broker/internal/session"

type Command interface {
	Name() string
	Description() string

	Arguments() []Argument

	CanBeCalled(*session.Session) error

	ValidateArgument(*session.Session, Argument, *string) error
	ParsePayload(*session.Session, map[string]*string) (error, interface{})
	OnCalled(*session.Session, interface{}) (error, interface{})

	GetResponse(*session.Session, interface{}) []string
}

type Argument struct {
	Key      string
	Name     string
	Required bool
	Help     string
}
