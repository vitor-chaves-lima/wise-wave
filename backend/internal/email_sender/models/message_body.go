package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TemplateName string

const (
	FirstAccessMagicLink TemplateName = "FirstAccessMagicLink"
)

type MessageBodyBase struct {
	RecipientEmail    string      `json:"recipientEmail"`
	EmailTemplateName string      `json:"emailTemplateName"`
	Data              interface{} `json:"data"`
}

func UnmarshallMessageBody(body string) (m MessageBodyBase, e error) {
	err := json.Unmarshal([]byte(body), &m)

	if err != nil {
		return m, errors.Join(fmt.Errorf("couldn't unmarshall message body"), err)
	}

	return m, nil
}
