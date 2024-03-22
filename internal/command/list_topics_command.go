package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

type ListTopicsCommand struct{}

func (listTopics *ListTopicsCommand) Name() string {
	return "list-topics"
}

func (listTopics *ListTopicsCommand) Description() string {
	return "List topics"
}

func (listTopics *ListTopicsCommand) Arguments() []Argument {
	return nil
}

func (listTopics *ListTopicsCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (listTopics *ListTopicsCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (listTopics *ListTopicsCommand) ParsePayload(*session.Session, map[string]*string) (error, interface{}) {
	return nil, nil
}

func (listTopics *ListTopicsCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	topics := subscription.ListTopics()

	return nil, topics
}

func (listTopics *ListTopicsCommand) GetResponse(sess *session.Session, payload interface{}) []string {
	topics := payload.([]*subscription.Topic)

	output := make([]string, len(topics))

	for i := range topics {
		topic := topics[i]
		subscribers := topic.Subscribers.([]*subscription.Subscriber)

		output[i] = fmt.Sprintf("Topic: %s (%d Subscribers)", topic.Name, len(subscribers))
	}

	return output
}
