package controller

import (
	"message_broker/internal/interpreter"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

func subscribeTopic(session *session.Session, payload interpreter.SubscriptionPayload) (error, *subscription.Topic) {
	err, topic := subscription.SubscribeToTopic(payload.Topic, session)

	return err, topic
}

func unsubscribeTopic(session *session.Session, payload interpreter.SubscriptionPayload) (error, string) {
	err := subscription.UnsubscribeFromTopic(payload.Topic, session)

	return err, payload.Topic
}

func listTopicsFromSession(session *session.Session) []*subscription.Topic {
	topics := subscription.GetTopicsFromSession(session)

	return topics
}

func listSubscribers(payload interpreter.SubscriptionPayload) (error, []*subscription.Subscriber) {
	err, subscribers := subscription.GetSubscribers(payload.Topic)

	return err, subscribers
}

func listTopics() []*subscription.Topic {
	topics := subscription.ListTopics()

	return topics
}
