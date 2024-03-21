package interpreter

type SubscriptionPayload struct {
	Topic string
}

func parseSubscriptionPayload(payloadData []string) (error, *SubscriptionPayload) {
	err, topic := retrieveValue(payloadData, "topic")
	if err != nil {
		return err, nil
	}

	payload := SubscriptionPayload{
		Topic: *topic,
	}

	return nil, &payload
}
