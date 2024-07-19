package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	log "github.com/sirupsen/logrus"
	"wisewave.tech/email_sender_service/internal/adapters"
	"wisewave.tech/email_sender_service/internal/application/usecases"
	"wisewave.tech/email_sender_service/internal/ports"
)

var (
	sesClient            *ses.Client
	ssmClient            *ssm.Client
	sesEmailer           ports.Emailer
	sqsConsumer          ports.QueueMessageConsumer
	sendHTMLEmailUseCase *usecases.SendEmailUseCase
)

func handle(ctx context.Context, event events.SQSEvent) {
	log.Info("Starting lambda function")
	err := sqsConsumer.Consume(event)
	if err != nil {
		panic(err)
	}
}

func getEmailSenderIdentityParameter() (emailSenderIdentityParameter string, err error) {
	emailSenderIdentityParameter = os.Getenv("SENDER_IDENTITY_PARAMETER")
	if emailSenderIdentityParameter == "" {
		panic("SENDER_IDENTITY_PARAMETER env var is not set!")
	}

	log.WithField("emailSenderIdentityParameter", emailSenderIdentityParameter).Info("Searching email sender identity parameter in SSM")
	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(emailSenderIdentityParameter),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", errors.Join(errors.New("couldn't find email sender identity parameter in SSM"), err)
	}

	return *param.Parameter.Value, nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	log.Info("Initializing SES client")
	sesClient = ses.NewFromConfig(cfg)

	log.Info("Initializing SSM client")
	ssmClient = ssm.NewFromConfig(cfg)

	emailSenderIdentity, err := getEmailSenderIdentityParameter()
	if err != nil {
		panic(err)
	}
	log.Println("Sender email identity", emailSenderIdentity)

	log.Info("Initializing SES emailer")
	sesEmailer = adapters.NewSESEmailer(sesClient, emailSenderIdentity)

	log.Info("Initializing send email use case")
	sendHTMLEmailUseCase, err = usecases.NewSendEmailUseCase(sesEmailer)
	if err != nil {
		panic(err)
	}

	log.Info("Initializing SQS queue message consumer")
	sqsConsumer = adapters.NewSQSQueueMessageConsumer(sendHTMLEmailUseCase)
}

func main() {
	lambda.Start(handle)
}
