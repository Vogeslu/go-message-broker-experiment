package interpreter

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"message_broker/internal/messagetype"
	"message_broker/internal/session"
	"strings"
)

func ParseMessageType(session *session.Session, lines []string) (error, messagetype.MessageType) {
	action := strings.TrimRight(lines[0], "\n")

	messageType := messagetype.GetMessageType(action)

	logger.Logger.Debug().Msg(fmt.Sprintf("Received action \"%s\" %s", messageType.String(), session.Conn.RemoteAddr().String()))

	if messageType == messagetype.Invalid {
		return errors.New("Invalid action"), messageType
	}

	return nil, messageType
}

func ParsePayload(payloadData []string, messageType messagetype.MessageType) (error, interface{}) {
	switch messageType {
	case messagetype.Greet:
		return parseGreetPayload(payloadData)
	case messagetype.SubscribeTopic,
		messagetype.UnsubscribeTopic,
		messagetype.ListSubscribers:
		return parseSubscriptionPayload(payloadData)
	case messagetype.SendMessage:
		return parseMessagePayload(payloadData)
	case messagetype.AcknowledgeMessage,
		messagetype.ListAcknowledgements:
		return parseMessageIdPayload(payloadData)
	case messagetype.Invalid:
		return errors.New("Invalid message type"), nil
	}

	return nil, nil
}
