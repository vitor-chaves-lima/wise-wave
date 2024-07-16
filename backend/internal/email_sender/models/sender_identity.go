package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SenderIdentity struct {
	ARN   string `json:"arn"`
	Email string `json:"email"`
}

func UnmarshallSenderIdentity(body string) (s SenderIdentity, e error) {
	err := json.Unmarshal([]byte(body), &s)

	if err != nil {
		return s, errors.Join(fmt.Errorf("couldn't unmarshall sender identity data"), err)
	}

	return s, nil
}
