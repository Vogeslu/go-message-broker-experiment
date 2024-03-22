package command

import (
	"errors"
	"message_broker/internal/messaging"
	"message_broker/internal/session"

	"github.com/google/uuid"
)

type ListAcknowledgementsCommand struct{}

func (listAcknowledgements *ListAcknowledgementsCommand) Name() string {
	return "list-acks"
}

func (listAcknowledgements *ListAcknowledgementsCommand) Description() string {
	return "List acknowledgements of a message"
}

func (listAcknowledgements *ListAcknowledgementsCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "id",
			Name:     "ID",
			Help:     "",
			Required: true,
		},
	}
}

func (listAcknowledgements *ListAcknowledgementsCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (listAcknowledgements *ListAcknowledgementsCommand) ValidateArgument(session *session.Session, argument Argument, value *string) error {
	if argument.Key == "id" {
		_, err := uuid.Parse(*value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (listAcknowledgements *ListAcknowledgementsCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	messageId, err := uuid.Parse(*data["id"])
	if err != nil {
		return err, nil
	}

	return nil, messageId
}

func (listAcknowledgements *ListAcknowledgementsCommand) OnCalled(sess *session.Session, payload interface{}) (error, interface{}) {
	messageId := payload.(uuid.UUID)

	output := make([]*messaging.SessionMessageAcknowledgement, 0)

	sessions := session.GetActiveSessions()
	acknowledgements := messaging.GetAcknowledgements(messageId)

	for i := range sessions {
		session := sessions[i]
		var foundAcknowledgement *messaging.MessageAcknowledgement = nil

		for j := range acknowledgements {
			acknowledgement := acknowledgements[j]

			if acknowledgement.Session == session {
				foundAcknowledgement = acknowledgement
				break
			}
		}

		sessionMessageAck := &messaging.SessionMessageAcknowledgement{
			Session:         session,
			Acknowledgement: foundAcknowledgement,
		}

		output = append(output, sessionMessageAck)
	}

	return nil, output
}

func (listAcknowledgements *ListAcknowledgementsCommand) GetResponse(session *session.Session, payload interface{}) []string {
	acknowledgements := payload.([]*messaging.SessionMessageAcknowledgement)

	output := make([]string, len(acknowledgements))

	for i := range acknowledgements {
		acknowledgement := acknowledgements[i]
		output[i] = acknowledgement.ToString()
	}

	return output
}
