package listener

import (
	"fmt"
	"io"
	"message_broker/internal/command"
	"message_broker/internal/config"
	"message_broker/internal/logger"
	"message_broker/internal/session"
	"message_broker/internal/subscription"
	"net"
	"strings"
)

func StartListener() {
	server := config.AppConfig.Server

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		logger.Logger.Err(err)
		return
	}

	defer listener.Close()

	logger.Logger.Info().Msg(fmt.Sprintf("Listener running on %s:%d", server.Host, server.Port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Logger.Err(err).Msg("Failed accepting connection")
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	sess := session.CreateSession(conn)

	buffer := make([]byte, 1024)
	var nextArgument *command.Argument = nil

	for {
		conn.Write([]byte(getSessionPrefix(sess, nextArgument)))
		nextArgument = nil

		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.Logger.Debug().AnErr("Connection read error", err)
			}

			return
		}

		line := strings.TrimRight(string(buffer[:n]), "\n")

		err, missingArguments, response := command.ParseLineInput(sess, line)
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: %s\n", err.Error())))

			if missingArguments == nil {
				continue
			}
		}

		if len(missingArguments) > 0 {
			nextArgument = &missingArguments[0]
			continue
		}

		for i := range response {
			conn.Write([]byte(response[i] + "\n"))
		}
	}

	subscription.UnsubscribeFromAllTopics(sess)
	session.RemoveSession(conn)
}

func getSessionPrefix(session *session.Session, nextArgument *command.Argument) string {
	prefix := ""

	if session.Name != nil {
		prefix = fmt.Sprintf("[%s] ", *session.Name)
	} else {
		prefix = "[?] "
	}

	if nextArgument != nil {
		prefix += fmt.Sprintf("%s: ", nextArgument.Name)
	} else {
		prefix += "> "
	}

	return prefix
}
