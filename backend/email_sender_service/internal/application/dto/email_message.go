package dto

import (
	"errors"
	"fmt"

	"wisewave.tech/email_sender_service/internal/domain"
)

var EmailTypeStrings = []string{
	"newUserMagicLink",
	"magicLink",
}

type EmailMessageDTO struct {
	EmailType string                 `json:"emailType"`
	To        string                 `json:"to"`
	Data      map[string]interface{} `json:"data"`
}

func getEmailTypeFromString(emailTypeString string) (emailType domain.EmailType, err error) {
	switch emailTypeString {
	case "newUserMagicLink":
		return domain.NewUserMagicLink, nil
	case "magicLink":
		return domain.MagicLink, nil
	default:
		return 0, fmt.Errorf("invalid email type")
	}
}

func NewEmailTemplateDataFromDTO(emailMessageDto EmailMessageDTO) (emailTemplateData domain.EmailTemplateData, err error) {
	emailType, err := getEmailTypeFromString(emailMessageDto.EmailType)
	if err != nil {
		return domain.EmailTemplateData{}, errors.Join(fmt.Errorf("couldn't isntantiate EmailTemplateData"), err)
	}

	return domain.EmailTemplateData{
		Type: emailType,
		Data: emailMessageDto.Data,
	}, nil
}
