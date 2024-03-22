package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

type ListSubscribersCommand struct{}

func (listSubscribers *ListSubscribersCommand) Name() string {
	return "list-subscribers"
}

func (listSubscribers *ListSubscribersCommand) Description() string {
	return "List subscribers of a topic"
}

func (listSubscribers *ListSubscribersCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "topic",
			Name:     "Topic",
			Help:     "",
			Required: true,
		},
	}
}

func (listSubscribers *ListSubscribersCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (listSubscribers *ListSubscribersCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (listSubscribers *ListSubscribersCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	return nil, data["topic"]
}

func (listSubscribers *ListSubscribersCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	topicName := payload.(*string)
	err, subscribers := subscription.GetSubscribers(*topicName)

	return err, subscribers
}

func (listSubscribers *ListSubscribersCommand) GetResponse(session *session.Session, payload interface{}) []string {
	subscribers := payload.([]*subscription.Subscriber)

	output := make([]string, len(subscribers))

	for i := range subscribers {
		session := subscribers[i].Session
		output[i] = fmt.Sprintf("Name: %s, Address: %s, Leader: %t", *session.Name, session.Conn.RemoteAddr().String(), subscribers[i].IsLeader)
	}

	return output
}
