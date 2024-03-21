package subscription

import "message_broker/internal/session"

type Subscriber struct {
	Session  *session.Session
	IsLeader bool
}

type Topic struct {
	Name        string
	Subscribers interface{}
}
