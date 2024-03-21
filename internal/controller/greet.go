package controller

import (
	"errors"
	"fmt"
	"message_broker/internal/interpreter"
	"message_broker/internal/logger"
	"message_broker/internal/session"
)

func greet(session *session.Session, payload interpreter.GreetPayload) error {
	if session.Greeted {
		return errors.New("Connection already greeted")
	}

	session.Name = &payload.Name
	session.Greeted = true

	logger.Logger.Info().Msg(fmt.Sprintf("Session %s greeted with name %s", session.Conn.RemoteAddr().String(), *session.Name))

	return nil
}
