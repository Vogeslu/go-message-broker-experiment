package command

import (
	"errors"
	"message_broker/internal/messaging"
	"message_broker/internal/session"

	"github.com/google/uuid"
)

type AcknowledgeMessageCommand struct{}

func (acknowledgeMessage *AcknowledgeMessageCommand) Name() string {
	return "ack-message"
}

func (acknowledgeMessage *AcknowledgeMessageCommand) Description() string {
	return "Acknowledge a message"
}

func (acknowledgeMessage *AcknowledgeMessageCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "id",
			Name:     "ID",
			Help:     "",
			Required: true,
		},
	}
}

func (acknowledgeMessage *AcknowledgeMessageCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (acknowledgeMessage *AcknowledgeMessageCommand) ValidateArgument(session *session.Session, argument Argument, value *string) error {
	if argument.Key == "id" {
		_, err := uuid.Parse(*value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (acknowledgeMessage *AcknowledgeMessageCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	messageId, err := uuid.Parse(*data["id"])
	if err != nil {
		return err, nil
	}

	return nil, messageId
}

func (acknowledgeMessage *AcknowledgeMessageCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	messageId := payload.(uuid.UUID)
	err, acknowledgement := messaging.AcknowledgeMessage(session, messageId)

	return err, acknowledgement
}

func (acknowledgeMessage *AcknowledgeMessageCommand) GetResponse(session *session.Session, payload interface{}) []string {
	return []string{"Message acknowledged"}
}
