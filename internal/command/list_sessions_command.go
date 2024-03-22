package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
)

type ListSessionsCommand struct{}

func (listSessions *ListSessionsCommand) Name() string {
	return "list-sessions"
}

func (listSessions *ListSessionsCommand) Description() string {
	return "List connected sessions"
}

func (listSessions *ListSessionsCommand) Arguments() []Argument {
	return nil
}

func (listSessions *ListSessionsCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (listSessions *ListSessionsCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (listSessions *ListSessionsCommand) ParsePayload(*session.Session, map[string]*string) (error, interface{}) {
	return nil, nil
}

func (listSessions *ListSessionsCommand) OnCalled(*session.Session, interface{}) (error, interface{}) {
	sessions := session.GetActiveSessions()

	return nil, sessions
}

func (listSessions *ListSessionsCommand) GetResponse(sess *session.Session, payload interface{}) []string {
	sessions := payload.([]*session.Session)

	output := make([]string, len(sessions))

	for i := range sessions {
		session := sessions[i]
		output[i] = fmt.Sprintf("Name: %s, Address: %s", *session.Name, session.Conn.RemoteAddr().String())
	}

	return output
}
