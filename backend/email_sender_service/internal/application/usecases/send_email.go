package usecases

import (
	"encoding/json"
	"errors"
	"fmt"

	"wisewave.tech/email_sender_service/internal/application/dto"
	"wisewave.tech/email_sender_service/internal/application/validators"
	"wisewave.tech/email_sender_service/internal/domain"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SendEmailUseCase struct {
	templateManager ports.TemplateManager
	emailer         ports.Emailer
}

func (u *SendEmailUseCase) Execute(messageBody string) (err error) {
	var emailMessageDTO dto.EmailMessageDTO
	if err := json.Unmarshal([]byte(messageBody), &emailMessageDTO); err != nil {
		return errors.Join(fmt.Errorf("couldn't unmarshal message body"), err)
	}

	if err := validators.ValidateEmailMessageDTO(emailMessageDTO); err != nil {
		return errors.Join(fmt.Errorf("invalid message format"), err)
	}

	emailType, err := domain.GetEmailTypeFromString(emailMessageDTO.EmailType)
	if err != nil {
		return errors.Join(fmt.Errorf("couldn't get email type"), err)
	}

	subject, body, err := u.templateManager.FormatEmail(emailType, emailMessageDTO.Data)
	if err != nil {
		return errors.Join(errors.New("couldn't format email"), err)
	}

	err = u.emailer.SendHTMLEmail(emailMessageDTO.To, subject, body)
	if err != nil {
		return errors.Join(errors.New("couldn't send email"), err)
	}

	return nil
}

func NewSendEmailUseCase(templateManager ports.TemplateManager, emailer ports.Emailer) *SendEmailUseCase {
	return &SendEmailUseCase{
		templateManager,
		emailer,
	}
}
