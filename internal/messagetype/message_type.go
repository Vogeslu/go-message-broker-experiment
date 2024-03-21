package messagetype

type MessageType int

const (
	Greet MessageType = iota
	ListSessions
	Ping

	SubscribeTopic
	UnsubscribeTopic
	ListSubscriptions
	ListSubscribers
	ListTopics

	SendMessage
	AcknowledgeMessage
	ListAcknowledgements

	Invalid
)

func (messageType MessageType) String() string {
	return []string{
		"greet",
		"list-sessions",
		"ping",

		"subscribe-topic",
		"unsubscribe-topic",
		"list-subscriptions",
		"list-subscribers",
		"list-topics",

		"send-message",
		"ack-message",
		"list-acks",

		"invalid",
	}[messageType]
}

func GetMessageType(data string) MessageType {
	for i := Greet; i < Invalid; i++ {
		messageType := MessageType(i)

		if messageType.String() == data {
			return messageType
		}
	}

	return Invalid
}
