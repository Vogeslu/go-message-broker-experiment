package listener

import (
	"fmt"
	"io"
	"message_broker/internal/config"
	"message_broker/internal/controller"
	"message_broker/internal/interpreter"
	"message_broker/internal/logger"
	"message_broker/internal/middleware"
	"message_broker/internal/output"
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
	var currentPayload []string

	for {
		conn.Write([]byte(getSessionPrefix(sess)))

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

		if line != "." {
			currentPayload = append(currentPayload, line)
		} else if len(currentPayload) > 0 {
			err, messageType := interpreter.ParseMessageType(sess, currentPayload)
			payloadData := currentPayload[1:]
			currentPayload = nil

			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			err = middleware.HandleMiddleware(sess, messageType)

			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			err, payload := interpreter.ParsePayload(payloadData, messageType)

			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			err, data := controller.HandleAction(sess, messageType, payload)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			output := output.OutputData(messageType, data)

			if output != nil {
				for i := range output {
					conn.Write([]byte(output[i] + "\n"))
				}
			}
		}
	}

	subscription.UnsubscribeFromAllTopics(sess)
	session.RemoveSession(conn)
}

func getSessionPrefix(session *session.Session) string {
	if session.Name != nil {
		return fmt.Sprintf("%s > ", *session.Name)
	} else {
		return "? > "
	}
}
