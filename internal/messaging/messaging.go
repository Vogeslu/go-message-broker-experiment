package messaging

import (
	"message_broker/internal/session"
	"message_broker/internal/subscription"
	"net"
	"strings"
	"time"

	"github.com/google/uuid"
)

var messages = make(map[string][]*Message)

func SendMessageToSubscribers(sender *session.Session, topic string, content string, manualAcknowledgementRequired bool) (error, *Message) {
	err, subscribers := subscription.GetSubscribers(topic)
	if err != nil {
		return err, nil
	}

	message := createMessage(sender, topic, content, manualAcknowledgementRequired)

	composedMessage := message.composeMessage(topic)

	payload := make([][]byte, len(composedMessage))
	for i := range composedMessage {
		payload[i] = []byte(composedMessage[i] + "\n")
	}

	for i := range subscribers {
		session := subscribers[i].Session

		if session == sender {
			continue
		}

		conn := session.Conn

		sendPayload(payload, conn)

		if !manualAcknowledgementRequired {
			AcknowledgeMessage(session, message.Id)
		}
	}

	return nil, message
}

func createMessage(sender *session.Session, topic string, content string, manualAcknowledgementRequired bool) *Message {
	message := &Message{
		Id:        uuid.New(),
		Content:   content,
		Sender:    sender,
		CreatedAt: time.Now(),
		Flags: MessageFlags{
			ManualAcknowledgementRequired: manualAcknowledgementRequired,
		},
	}

	messages[topic] = append(messages[topic], message)

	return message
}

func sendPayload(payload [][]byte, conn net.Conn) {
	for i := range payload {
		conn.Write(payload[i])
	}
}

func (message Message) composeMessage(topic string) []string {
	output := []string{
		message.Id.String(),
		topic,
		*message.Sender.Name,
		message.Content,
	}

	flags := message.Flags.getFlags()

	if len(flags) > 0 {
		output = append(output, flags)
	}

	return output
}

func (messageFlags MessageFlags) getFlags() string {
	parts := make([]string, 0)

	if messageFlags.ManualAcknowledgementRequired {
		parts = append(parts, "mack")
	}

	return strings.Join(parts, ";")
}
