package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

type UnsubscribeTopicCommand struct{}

func (unsubscribeTopic *UnsubscribeTopicCommand) Name() string {
	return "unsubscribe-topic"
}

func (unsubscribeTopic *UnsubscribeTopicCommand) Description() string {
	return "Unsubscribe a topic"
}

func (unsubscribeTopic *UnsubscribeTopicCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "topic",
			Name:     "Topic",
			Help:     "",
			Required: true,
		},
	}
}

func (unsubscribeTopic *UnsubscribeTopicCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (unsubscribeTopic *UnsubscribeTopicCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (unsubscribeTopic *UnsubscribeTopicCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	return nil, data["topic"]
}

func (unsubscribeTopic *UnsubscribeTopicCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	topicName := payload.(*string)
	err := subscription.UnsubscribeFromTopic(*topicName, session)

	return err, topicName
}

func (unsubscribeTopic *UnsubscribeTopicCommand) GetResponse(session *session.Session, payload interface{}) []string {
	topicName := payload.(*string)
	return []string{fmt.Sprintf("Unsubscribed topic %s", *topicName)}
}
