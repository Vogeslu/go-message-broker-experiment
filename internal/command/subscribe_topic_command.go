package command

import (
	"errors"
	"fmt"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

type SubscribeTopicCommand struct{}

func (subscribeTopic *SubscribeTopicCommand) Name() string {
	return "subscribe-topic"
}

func (subscribeTopic *SubscribeTopicCommand) Description() string {
	return "Subscribe a topic"
}

func (subscribeTopic *SubscribeTopicCommand) Arguments() []Argument {
	return []Argument{
		{
			Key:      "topic",
			Name:     "Topic",
			Help:     "",
			Required: true,
		},
	}
}

func (subscribeTopic *SubscribeTopicCommand) CanBeCalled(session *session.Session) error {
	if !session.Greeted {
		return errors.New("Session not greeted yet")
	}

	return nil
}

func (subscribeTopic *SubscribeTopicCommand) ValidateArgument(*session.Session, Argument, *string) error {
	return nil
}

func (subscribeTopic *SubscribeTopicCommand) ParsePayload(session *session.Session, data map[string]*string) (error, interface{}) {
	return nil, data["topic"]
}

func (subscribeTopic *SubscribeTopicCommand) OnCalled(session *session.Session, payload interface{}) (error, interface{}) {
	topicName := payload.(*string)
	err, topic := subscription.SubscribeToTopic(*topicName, session)

	return err, topic
}

func (subscribeTopic *SubscribeTopicCommand) GetResponse(session *session.Session, payload interface{}) []string {
	topic := payload.(*subscription.Topic)
	return []string{fmt.Sprintf("Subscribed topic %s", topic.Name)}
}
