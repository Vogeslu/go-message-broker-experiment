package interpreter

import "github.com/google/uuid"

type MessagePayload struct {
	Topic   string
	Message string
	Flags   MessagePayloadFlags
}

type MessagePayloadFlags struct {
	ManualAcknowledgementRequired bool
}

type MessageIdPayload struct {
	Id uuid.UUID
}

func parseMessagePayload(payloadData []string) (error, *MessagePayload) {
	err, topic := retrieveValue(payloadData, "topic")
	if err != nil {
		return err, nil
	}

	err, message := retrieveValue(payloadData, "message")
	if err != nil {
		return err, nil
	}

	var manualAcknowledgementRequired = false
	_, manualAck := retrieveValue(payloadData, "manualAcc")
	if manualAck != nil {
		manualAcknowledgementRequired = *manualAck == "true"
	}

	payload := MessagePayload{
		Topic:   *topic,
		Message: *message,
		Flags: MessagePayloadFlags{
			ManualAcknowledgementRequired: manualAcknowledgementRequired,
		},
	}

	return nil, &payload
}

func parseMessageIdPayload(payloadData []string) (error, *MessageIdPayload) {
	err, rawId := retrieveValue(payloadData, "id")
	if err != nil {
		return err, nil
	}

	uuid, err := uuid.Parse(*rawId)
	if err != nil {
		return err, nil
	}

	payload := MessageIdPayload{
		Id: uuid,
	}

	return nil, &payload
}
