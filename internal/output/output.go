package output

import (
	"fmt"
	"message_broker/internal/messagetype"
	"message_broker/internal/messaging"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
)

func OutputData(messageType messagetype.MessageType, data interface{}) []string {
	switch messageType {
	case messagetype.Greet:
		return []string{fmt.Sprintf("Greeted")}
	case messagetype.ListSessions:
		sessions := data.([]*session.Session)

		output := make([]string, len(sessions))

		for i := range sessions {
			session := sessions[i]
			output[i] = fmt.Sprintf("Name: %s, Address: %s", *session.Name, session.Conn.RemoteAddr().String())
		}

		return output
	case messagetype.Ping:
		return []string{"Pong"}
	case messagetype.SubscribeTopic:
		topic := data.(*subscription.Topic)

		return []string{fmt.Sprintf("Subscribed topic %s", topic.Name)}
	case messagetype.UnsubscribeTopic:
		topic := data.(string)

		return []string{fmt.Sprintf("Unsubscribed topic %s", topic)}
	case messagetype.ListSubscriptions, messagetype.ListTopics:
		topics := data.([]*subscription.Topic)

		output := make([]string, len(topics))

		for i := range topics {
			topic := topics[i]
			subscribers := topic.Subscribers.([]*subscription.Subscriber)

			output[i] = fmt.Sprintf("Topic: %s (%d Subscribers)", topic.Name, len(subscribers))
		}

		return output
	case messagetype.ListSubscribers:
		subscribers := data.([]*subscription.Subscriber)

		output := make([]string, len(subscribers))

		for i := range subscribers {
			session := subscribers[i].Session
			output[i] = fmt.Sprintf("Name: %s, Address: %s, Leader: %t", *session.Name, session.Conn.RemoteAddr().String(), subscribers[i].IsLeader)
		}

		return output
	case messagetype.SendMessage:
		message := data.(*messaging.Message)

		return []string{message.Id.String()}
	case messagetype.AcknowledgeMessage:
		return []string{"Message acknowledged"}
	case messagetype.ListAcknowledgements:
		acknowledgements := data.([]*messaging.SessionMessageAcknowledgement)

		output := make([]string, len(acknowledgements))

		for i := range acknowledgements {
			acknowledgement := acknowledgements[i]
			output[i] = acknowledgement.ToString()
		}

		return output
	}

	return nil
}
