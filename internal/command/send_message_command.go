package command

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"message_broker/internal/messaging"
	"message_broker/internal/session"
)

type SendMessageCommand struct{}

type messagePayload struct {
	Topic   string
	Message string
	Flags   messagePayloadFlags
}

type messagePayloadFlags struct {
	ManualAcknowledgementRequired bool
}

func (sendMessage *SendMessageCommand) Name() string {
	return "send-message"
}

func (sendMessage *SendMessageCommand) Description() string {
	return "Send message to a topic"
}

func (sendMessage *SendMessageCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "topic",
			Name:     "Topic",
			Help:     "",
			Required: true,
		},
		{
			Key:      "message",
			Name:     "Message",
			Help:     "",
			Required: true,
		},
		{
			Key:      "manualAcc",
			Name:     "Manual acknowledgement",
			Help:     "",
			Required: false,
		},
	}
}

func (sendMessage *SendMessageCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (sendMessage *SendMessageCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (sendMessage *SendMessageCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	payload := &messagePayload{
		Topic:   *data["topic"],
		Message: *data["message"],
		Flags: messagePayloadFlags{
			ManualAcknowledgementRequired: data["manualAcc"] != nil && *data["manualAcc"] == "true",
		},
	}

	return nil, payload
}

func (sendMessage *SendMessageCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	messagePayload := payload.(*messagePayload)

	logger.Logger.Info().Msg(fmt.Sprintf("Session %s is sending message to topic %s", session.Conn.RemoteAddr().String(), messagePayload.Topic))

	err, message := messaging.SendMessageToSubscribers(session, messagePayload.Topic, messagePayload.Message, messagePayload.Flags.ManualAcknowledgementRequired)

	return err, message
}

func (sendMessage *SendMessageCommand) GetResponse(session *session.Session, payload interface{}) []string {
	message := payload.(*messaging.Message)
	return []string{message.Id.String()}
}
