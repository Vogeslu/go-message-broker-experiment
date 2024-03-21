package interpreter

import (
	"errors"
	"fmt"
	"strings"
)

func retrieveValue(lines []string, key string) (error, *string) {
	search := fmt.Sprintf("%s:", key)

	for i := range lines {
		if strings.HasPrefix(lines[i], search) {
			value := strings.TrimSpace(strings.Join(strings.Split(lines[i], search)[1:], ""))
			return nil, &value
		}
	}

	return errors.New(fmt.Sprintf("Value for key %s not found", key)), nil
}
