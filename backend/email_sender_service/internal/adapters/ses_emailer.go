package adapters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SESEmailer struct {
	ctx    context.Context
	client *ses.Client
	source string
}

func NewSESEmailer(ctx context.Context, client *ses.Client, source string) ports.Emailer {
	fields := logrus.Fields{
		"type":          "adapter",
		"name":          "ses_emailer",
		"function_name": "NewSESEmailer",
	}
	logger := lib.LoggerFromContext(ctx).WithFields(fields)

	logger.Info("creating SESEmailer")
	return &SESEmailer{ctx, client, source}
}

func (e *SESEmailer) SendHTMLEmail(to, subject, htmlBody string) error {
	fields := logrus.Fields{
		"type":          "adapter",
		"name":          "ses_emailer",
		"function_name": "emailer.SendHTMLEmail",
		"to":            to,
		"subject":       subject,
	}
	logger := lib.LoggerFromContext(e.ctx).WithFields(fields)

	logger.Info("creating send email request input")
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Data: &htmlBody,
				},
			},
			Subject: &types.Content{
				Data: &subject,
			},
		},
		Source: &e.source,
	}

	logger.Info("sending email request input")
	_, err := e.client.SendEmail(context.Background(), input)
	return err
}
