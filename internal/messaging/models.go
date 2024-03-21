package messaging

import (
	"fmt"
	"message_broker/internal/session"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	Content   string
	Sender    *session.Session
	CreatedAt time.Time
	Flags     MessageFlags
}

type MessageFlags struct {
	ManualAcknowledgementRequired bool
}

type MessageAcknowledgement struct {
	Session    *session.Session
	ReceivedAt time.Time
}

type SessionMessageAcknowledgement struct {
	Session         *session.Session
	Acknowledgement *MessageAcknowledgement
}

func (acknowledgement *SessionMessageAcknowledgement) ToString() string {
	parts := []string{
		fmt.Sprintf("Name: %s", *acknowledgement.Session.Name),
		fmt.Sprintf("Received: %t", acknowledgement.Acknowledgement != nil),
	}

	if acknowledgement.Acknowledgement != nil {
		parts = append(parts, fmt.Sprintf("ReceivedAt: %s", acknowledgement.Acknowledgement.ReceivedAt.String()))
	}

	return strings.Join(parts, ", ")
}
