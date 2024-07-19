package adapters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SESEmailer struct {
	client *ses.Client
	source string
}

func NewSESEmailer(ctx context.Context, client *ses.Client, source string) ports.Emailer {
	return &SESEmailer{client: client, source: source}
}

func (e *SESEmailer) SendHTMLEmail(to, subject, htmlBody string) error {
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

	_, err := e.client.SendEmail(context.Background(), input)
	return err
}
