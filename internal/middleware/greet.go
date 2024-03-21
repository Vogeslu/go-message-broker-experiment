package middleware

import (
	"errors"
	"message_broker/internal/messagetype"
	"message_broker/internal/session"
)

func HandleGreet(session *session.Session, messageType messagetype.MessageType) error {
	if messageType == messagetype.Greet {
		return nil
	}

	if !session.Greeted {
		return errors.New("Session has not been greeted yet")
	}

	return nil
}
