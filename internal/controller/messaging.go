package controller

import (
	"fmt"
	"message_broker/internal/interpreter"
	"message_broker/internal/logger"
	"message_broker/internal/messaging"
	"message_broker/internal/session"
)

func sendMessage(session *session.Session, payload interpreter.MessagePayload) (error, *messaging.Message) {
	logger.Logger.Info().Msg(fmt.Sprintf("Session %s is sending message to topic %s", session.Conn.RemoteAddr().String(), payload.Topic))

	err, message := messaging.SendMessageToSubscribers(session, payload.Topic, payload.Message, payload.Flags.ManualAcknowledgementRequired)

	return err, message
}

func acknowledgeMessage(session *session.Session, messagePayload interpreter.MessageIdPayload) (error, *messaging.MessageAcknowledgement) {
	err, acknowledgement := messaging.AcknowledgeMessage(session, messagePayload.Id)

	return err, acknowledgement
}

func listAcknowledgements(messagePayload interpreter.MessageIdPayload) []*messaging.SessionMessageAcknowledgement {
	output := make([]*messaging.SessionMessageAcknowledgement, 0)

	sessions := session.GetActiveSessions()
	acknowledgements := messaging.GetAcknowledgements(messagePayload.Id)

	for i := range sessions {
		session := sessions[i]
		var foundAcknowledgement *messaging.MessageAcknowledgement = nil

		for j := range acknowledgements {
			acknowledgement := acknowledgements[j]

			if acknowledgement.Session == session {
				foundAcknowledgement = acknowledgement
				break
			}
		}

		sessionMessageAck := &messaging.SessionMessageAcknowledgement{
			Session:         session,
			Acknowledgement: foundAcknowledgement,
		}

		output = append(output, sessionMessageAck)
	}

	return output
}
