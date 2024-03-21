package session

import (
	"errors"
	"fmt"
	"message_broker/internal/logger"
	"net"
)

var activeSessions []*Session

func GetActiveSessions() []*Session {
	return activeSessions
}

func FindSession(conn net.Conn) (error, *Session) {
	for i := range activeSessions {
		if activeSessions[i].Conn == conn {
			return nil, activeSessions[i]
		}
	}

	return errors.New("Session not found"), nil
}

func FindIndex(conn net.Conn) (error, int) {
	for i := range activeSessions {
		if activeSessions[i].Conn == conn {
			return nil, i
		}
	}

	return errors.New("Index not found"), -1
}

func RemoveSession(conn net.Conn) {
	err, index := FindIndex(conn)

	if err == nil {
		activeSessions[index] = activeSessions[len(activeSessions)-1]
		activeSessions = activeSessions[:len(activeSessions)-1]

		address := conn.RemoteAddr().String()

		logger.Logger.Info().Msg(fmt.Sprintf("Session %s closed and removed", address))
	}
}

func CreateSession(conn net.Conn) *Session {
	session := Session{
		Conn:    conn,
		Greeted: false,
	}

	var sessionPtr *Session
	sessionPtr = &session

	activeSessions = append(activeSessions, sessionPtr)

	address := conn.RemoteAddr().String()

	logger.Logger.Info().Msg(fmt.Sprintf("Session %s opened", address))

	return sessionPtr
}
