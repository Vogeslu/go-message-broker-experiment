package interpreter

type GreetPayload struct {
	Name string
}

func parseGreetPayload(payloadData []string) (error, *GreetPayload) {
	err, name := retrieveValue(payloadData, "name")
	if err != nil {
		return err, nil
	}

	payload := GreetPayload{
		Name: *name,
	}

	return nil, &payload
}
