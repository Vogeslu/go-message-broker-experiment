package messaging

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"message_broker/internal/session"
	"time"

	"github.com/google/uuid"
)

var acknowledgements = make(map[uuid.UUID][]*MessageAcknowledgement)

func AcknowledgeMessage(session *session.Session, messageId uuid.UUID) (error, *MessageAcknowledgement) {
	if HasAcknowledgedMessage(session, messageId) {
		return errors.New("Message has already been acknowledged"), nil
	}

	acknowledgement := &MessageAcknowledgement{
		Session:    session,
		ReceivedAt: time.Now(),
	}

	acknowledgements[messageId] = append(acknowledgements[messageId], acknowledgement)

	logger.Logger.Debug().Msg(fmt.Sprintf("Message %s has been acknowledged by %s", messageId.String(), *session.Name))

	return nil, acknowledgement
}

func HasAcknowledgedMessage(session *session.Session, messageId uuid.UUID) bool {
	messageAcks, ok := acknowledgements[messageId]
	if ok {
		for i := range messageAcks {
			if messageAcks[i].Session == session {
				return true
			}
		}
	}

	return false
}

func GetAcknowledgements(messageId uuid.UUID) []*MessageAcknowledgement {
	messageAcks, ok := acknowledgements[messageId]
	if ok {
		return messageAcks
	}

	return []*MessageAcknowledgement{}
}
