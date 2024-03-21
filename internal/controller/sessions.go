package controller

import (
	"message_broker/internal/session"
)

func listSessions() []*session.Session {
	sessions := session.GetActiveSessions()

	return sessions
}
