package use_cases

import (
	"errors"

	"wisewave.tech/email_sender_service/internal/domain"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SendEmailUseCase struct {
	templateManager ports.TemplateManager
	emailer         ports.Emailer
}

func (u *SendEmailUseCase) Execute(emailTemplate domain.EmailTemplate, data map[string]interface{}, to string) (err error) {
	subject, body, err := u.templateManager.FormatEmail(domain.MagicLink, data)
	if err != nil {
		return errors.Join(errors.New("couldn't format email"), err)
	}

	err = u.emailer.SendHTMLEmail(to, subject, body)
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
