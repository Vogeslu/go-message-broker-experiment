package controller

import (
	"message_broker/internal/interpreter"
	"message_broker/internal/messagetype"
	"message_broker/internal/session"
)

func HandleAction(session *session.Session, messageType messagetype.MessageType, payload interface{}) (error, interface{}) {
	switch messageType {
	case messagetype.Greet:
		greetPayload := *payload.(*interpreter.GreetPayload)
		err := greet(session, greetPayload)
		if err != nil {
			return err, nil
		}
	case messagetype.ListSessions:
		sessions := listSessions()

		return nil, sessions
	case messagetype.SubscribeTopic:
		subscriptionPayload := *payload.(*interpreter.SubscriptionPayload)
		err, subscription := subscribeTopic(session, subscriptionPayload)

		return err, subscription
	case messagetype.UnsubscribeTopic:
		subscriptionPayload := *payload.(*interpreter.SubscriptionPayload)
		err, topic := unsubscribeTopic(session, subscriptionPayload)

		return err, topic
	case messagetype.ListSubscriptions:
		topics := listTopicsFromSession(session)

		return nil, topics
	case messagetype.ListSubscribers:
		subscriptionPayload := *payload.(*interpreter.SubscriptionPayload)
		err, subscribers := listSubscribers(subscriptionPayload)

		return err, subscribers
	case messagetype.ListTopics:
		topics := listTopics()

		return nil, topics
	case messagetype.SendMessage:
		messagePayload := *payload.(*interpreter.MessagePayload)
		err, message := sendMessage(session, messagePayload)

		return err, message
	case messagetype.AcknowledgeMessage:
		messagePayload := *payload.(*interpreter.MessageIdPayload)
		err, acknowledgement := acknowledgeMessage(session, messagePayload)

		return err, acknowledgement
	case messagetype.ListAcknowledgements:
		messagePayload := *payload.(*interpreter.MessageIdPayload)
		acknowledgements := listAcknowledgements(messagePayload)

		return nil, acknowledgements
	}

	return nil, nil
}
