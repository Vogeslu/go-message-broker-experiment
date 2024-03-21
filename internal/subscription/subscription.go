package subscription

import (
	"errors"
	"fmt"
	"math/rand"
	"message_broker/internal/logger"
	"message_broker/internal/session"
)

var topics []*Topic

func ListTopics() []*Topic {
	var output []*Topic

	for i := range topics {
		output = append(output, topics[i])
	}

	return output
}

func GetTopicsFromSession(session *session.Session) []*Topic {
	var output []*Topic

	for i := range topics {
		topic := topics[i]

		if HasSubscribedTopic(session, topic) {
			output = append(output, topic)
		}
	}

	return output
}

func HasSubscribedTopic(session *session.Session, topic *Topic) bool {
	subscribers := topic.Subscribers.([]*Subscriber)

	for i := range subscribers {
		if subscribers[i].Session == session {
			return true
		}
	}

	return false
}

func GetTopic(name string) (error, *Topic) {
	for i := range topics {
		topic := topics[i]

		if topic.Name == name {
			return nil, topic
		}
	}

	return errors.New("Topic not found"), nil
}

func GetOrCreateTopic(name string) *Topic {
	for i := range topics {
		topic := topics[i]

		if topic.Name == name {
			return topic
		}
	}

	topic := &Topic{
		Name:        name,
		Subscribers: []*Subscriber{},
	}

	topics = append(topics, topic)

	logger.Logger.Info().Msg(fmt.Sprintf("Topic %s has been created", topic.Name))

	return topic
}

func SubscribeToTopic(topicName string, session *session.Session) (error, *Topic) {
	topic := GetOrCreateTopic(topicName)

	if HasSubscribedTopic(session, topic) {
		return errors.New("Session already subscribed to topic"), nil
	}

	addSessionAsSubscriber(topic, session)

	logger.Logger.Info().Msg(fmt.Sprintf("Session %s subscribed topic %s", session.Conn.RemoteAddr().String(), topic.Name))

	return nil, topic
}

func addSessionAsSubscriber(topic *Topic, session *session.Session) {
	subscribers := topic.Subscribers.([]*Subscriber)
	isLeader := len(subscribers) == 0

	subscriber := &Subscriber{
		Session:  session,
		IsLeader: isLeader,
	}

	subscribers = append(subscribers, subscriber)
	topic.Subscribers = subscribers
}

func removeSessionAsSubscriber(topic *Topic, session *session.Session) {
	subscribers := topic.Subscribers.([]*Subscriber)
	var wasLeader = false
	var sessionFound = false

	for i := range subscribers {
		if subscribers[i].Session == session {
			wasLeader = subscribers[i].IsLeader
			sessionFound = true

			subscribers[i] = subscribers[len(subscribers)-1]
			subscribers = subscribers[:len(subscribers)-1]

			topic.Subscribers = subscribers

			logger.Logger.Info().Msg(fmt.Sprintf("Session %s unsubscribed topic %s", session.Conn.RemoteAddr().String(), topic.Name))

			break
		}
	}

	if sessionFound {
		if len(subscribers) == 0 {
			for i := range topics {
				if topics[i] == topic {
					topics[i] = topics[len(topics)-1]
					topics = topics[:len(topics)-1]

					logger.Logger.Info().Msg(fmt.Sprintf("Topic %s with zero subscribers has been deleted", topic.Name))

					break
				}
			}
		} else if wasLeader {
			newLeaderIndex := rand.Intn(len(subscribers))
			subscriber := subscribers[newLeaderIndex]

			subscriber.IsLeader = true

			logger.Logger.Info().Msg(fmt.Sprintf("Session %s has been assigned as new leader for topic %s", *subscriber.Session.Name, topic.Name))
		}
	}
}

func UnsubscribeFromTopic(topicName string, session *session.Session) error {
	topic := GetOrCreateTopic(topicName)

	if HasSubscribedTopic(session, topic) {

		removeSessionAsSubscriber(topic, session)

		return nil
	} else {
		return errors.New("Session not subscribed to topic")
	}
}

func UnsubscribeFromAllTopics(session *session.Session) {
	topics := GetTopicsFromSession(session)

	for i := range topics {
		removeSessionAsSubscriber(topics[i], session)
	}
}

func GetSubscribers(topicName string) (error, []*Subscriber) {
	err, topic := GetTopic(topicName)
	if err != nil {
		return err, nil
	}

	subscribers := topic.Subscribers.([]*Subscriber)

	return nil, subscribers
}
