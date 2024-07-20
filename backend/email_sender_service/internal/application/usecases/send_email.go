package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"wisewave.tech/email_sender_service/internal/application/dto"
	"wisewave.tech/email_sender_service/internal/application/managers"
	"wisewave.tech/email_sender_service/internal/application/validators"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SendEmailUseCase struct {
	ctx                  context.Context
	emailer              ports.Emailer
	emailTemplateManager *managers.EmailTemplateManager
}

func (u *SendEmailUseCase) Execute(messageBody string) (err error) {
	var emailMessageDTO dto.EmailMessageDTO
	if err := json.Unmarshal([]byte(messageBody), &emailMessageDTO); err != nil {
		return errors.Join(fmt.Errorf("couldn't unmarshal message body"), err)
	}

	if err := validators.ValidateEmailMessageDTO(emailMessageDTO); err != nil {
		return errors.Join(fmt.Errorf("invalid message format"), err)
	}

	emailTemplateData, err := dto.NewEmailTemplateDataFromDTO(emailMessageDTO)
	if err != nil {
		return errors.Join(fmt.Errorf("couldn't get email type"), err)
	}

	subject, body, err := u.emailTemplateManager.FormatEmail(emailTemplateData)
	if err != nil {
		return errors.Join(errors.New("couldn't format email"), err)
	}

	err = u.emailer.SendHTMLEmail(emailMessageDTO.To, subject, body)
	if err != nil {
		return errors.Join(errors.New("couldn't send email"), err)
	}

	return nil
}

func NewSendEmailUseCase(ctx context.Context, emailer ports.Emailer, emailTemplateManager *managers.EmailTemplateManager) (sendEmailUserCase *SendEmailUseCase, err error) {
	return &SendEmailUseCase{
		ctx,
		emailer,
		emailTemplateManager,
	}, nil
}
