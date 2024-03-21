package middleware

import (
	"message_broker/internal/messagetype"
	"message_broker/internal/session"
)

func HandleMiddleware(session *session.Session, messageType messagetype.MessageType) error {
	if err := HandleGreet(session, messageType); err != nil {
		return err
	}

	return nil
}
