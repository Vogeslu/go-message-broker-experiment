package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

type ListSubscriptionsCommand struct{}

func (listSubscriptions *ListSubscriptionsCommand) Name() string {
	return "list-subscriptions"
}

func (listSubscriptions *ListSubscriptionsCommand) Description() string {
	return "List subscribed topics"
}

func (listSubscriptions *ListSubscriptionsCommand) Arguments() []Argument {
	return nil
}

func (listSubscriptions *ListSubscriptionsCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (listSubscriptions *ListSubscriptionsCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (listSubscriptions *ListSubscriptionsCommand) ParsePayload(*session.Session, map[string]*string) (error, interface{}) {
	return nil, nil
}

func (listSubscriptions *ListSubscriptionsCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	topics := subscription.GetTopicsFromSession(session)

	return nil, topics
}

func (listSubscriptions *ListSubscriptionsCommand) GetResponse(sess *session.Session, payload interface{}) []string {
	topics := payload.([]*subscription.Topic)

	output := make([]string, len(topics))

	for i := range topics {
		topic := topics[i]
		subscribers := topic.Subscribers.([]*subscription.Subscriber)

		output[i] = fmt.Sprintf("Topic: %s (%d Subscribers)", topic.Name, len(subscribers))
	}

	return output
}
