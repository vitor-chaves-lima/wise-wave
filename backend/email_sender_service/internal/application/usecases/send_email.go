package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/email_sender_service/internal/application/dto"
	"wisewave.tech/email_sender_service/internal/application/managers"
	"wisewave.tech/email_sender_service/internal/application/validators"
	"wisewave.tech/email_sender_service/internal/ports"
)

var logBaseFields = logrus.Fields{
	"type": "usecase",
}

type SendEmailUseCase struct {
	ctx                  context.Context
	emailer              ports.Emailer
	emailTemplateManager *managers.EmailTemplateManager
}

func NewSendEmailUseCase(ctx context.Context, emailer ports.Emailer, emailTemplateManager *managers.EmailTemplateManager) (sendEmailUserCase *SendEmailUseCase, err error) {
	logger := lib.LoggerFromContext(ctx).WithFields(logBaseFields)

	logger.Info("creating SendEmailUseCase")
	return &SendEmailUseCase{
		ctx,
		emailer,
		emailTemplateManager,
	}, nil
}

func (u *SendEmailUseCase) Execute(messageBody string) (err error) {
	logger := lib.LoggerFromContext(u.ctx).WithFields(logBaseFields)

	logger.Info("creating email message dto")
	var emailMessageDTO dto.EmailMessageDTO
	if err := json.Unmarshal([]byte(messageBody), &emailMessageDTO); err != nil {
		err = errors.Join(fmt.Errorf("couldn't unmarshal message body"), err)
		logger.Error(err)
		return err
	}

	logger.Info("validating email message dto")
	if err := validators.ValidateEmailMessageDTO(emailMessageDTO); err != nil {
		err = errors.Join(fmt.Errorf("invalid message format"), err)
		logger.Error(err)
		return err
	}

	logger.Info("creating email template data")
	emailTemplateData, err := dto.NewEmailTemplateDataFromDTO(emailMessageDTO)
	if err != nil {
		err = errors.Join(fmt.Errorf("couldn't get email type"), err)
		logger.Error(err)
		return err
	}

	logger.Info("formatting email")
	subject, body, err := u.emailTemplateManager.FormatEmail(emailTemplateData)
	if err != nil {
		err = errors.Join(errors.New("couldn't format email"), err)
		logger.Error(err)
		return err
	}

	logger.Info("sending email")
	err = u.emailer.SendHTMLEmail(emailMessageDTO.To, subject, body)
	if err != nil {
		err = errors.Join(errors.New("couldn't send email"), err)
		logger.Error(err)
		return err
	}

	return nil
}
