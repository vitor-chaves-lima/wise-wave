package email_sender_service

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
)

type SQSEmailSenderServiceMessagePublisher struct {
	logger              *logrus.Entry
	sqsClient           *sqs.Client
	emailSenderQueueUrl string
}

func NewSQSEmailSenderServiceMessagePublisher(ctx context.Context, sqsClient *sqs.Client, emailSenderQueueUrl string) *SQSEmailSenderServiceMessagePublisher {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "adapter",
		"port": "sqs_email_sender_service_message_publisher",
	})

	return &SQSEmailSenderServiceMessagePublisher{
		logger,
		sqsClient,
		emailSenderQueueUrl,
	}
}

func (a *SQSEmailSenderServiceMessagePublisher) SendNewUserMagicLinkEmail(userEmail, magicLink string) error {
	logger := a.logger.WithFields(logrus.Fields{
		"userEmail": userEmail,
	})

	logger.Info("creating new user magic link email message body")
	messageBodyData := map[string]interface{}{
		"emailType": "newUserMagicLink",
		"to":        userEmail,
		"data": map[string]string{
			"link": magicLink,
		},
	}

	logger.Info("marshalling new user magic link email message body")
	messageBody, err := json.Marshal(messageBodyData)
	if err != nil {
		logger.WithError(err).Error("failed to marshal message body")
		return err
	}

	logger.Info("sending new user magic link email message")
	_, err = a.sqsClient.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageGroupId: aws.String(userEmail),
		QueueUrl:       aws.String(a.emailSenderQueueUrl),
		MessageBody:    aws.String(string(messageBody)),
	})
	if err != nil {
		logger.WithError(err).Error("failed to send message")
		return err
	}

	return nil
}

func (a *SQSEmailSenderServiceMessagePublisher) SendMagicLinkEmail(userEmail, magicLink string) error {
	logger := a.logger.WithFields(logrus.Fields{
		"userEmail": userEmail,
	})

	logger.Info("creating magic link email message body")
	messageBodyData := map[string]interface{}{
		"emailType": "magicLink",
		"to":        userEmail,
		"data": map[string]string{
			"link": magicLink,
		},
	}

	logger.Info("marshalling magic link email message body")
	messageBody, err := json.Marshal(messageBodyData)
	if err != nil {
		logger.WithError(err).Error("failed to marshal message body")
		return err
	}

	logger.Info("sending magic link email message")
	_, err = a.sqsClient.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageGroupId: aws.String(userEmail),
		QueueUrl:       aws.String(a.emailSenderQueueUrl),
		MessageBody:    aws.String(string(messageBody)),
	})
	if err != nil {
		logger.WithError(err).Error("failed to send message")
		return err
	}

	return nil
}
