package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
	"wisewave.tech/email_sender_service/internal/adapters"
	"wisewave.tech/email_sender_service/internal/application/usecases"
)

var (
	sesClient           *ses.Client
	ssmClient           *ssm.Client
	emailSenderIdentity string
)

func createLogger(ctx context.Context) *logrus.Entry {
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	contextFields := logrus.Fields{
		"requestId":          lambdaContext.AwsRequestID,
		"invokedFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := logrus.New().WithField("type", "lambda.handler").WithField("record", contextFields)
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func handler(ctx context.Context, event events.SQSEvent) {
	logger := createLogger(ctx)
	logger.Info("Starting lambda function handler")

	logger.Info("Initializing SES emailer")
	sesEmailer := adapters.NewSESEmailer(ctx, sesClient, emailSenderIdentity)

	logger.Info("Initializing send email usecase")
	sendHTMLEmailUseCase, err := usecases.NewSendEmailUseCase(sesEmailer)
	if err != nil {
		panic(err)
	}

	logger.Info("Initializing SQS queue message consumer")
	sqsConsumer := adapters.NewSQSQueueMessageConsumer(sendHTMLEmailUseCase)

	err = sqsConsumer.Consume(event)
	if err != nil {
		panic(err)
	}
}

func getEmailSenderIdentityParameter(logger *logrus.Entry) (emailSenderIdentityParameter string, err error) {
	emailSenderIdentityParameter = os.Getenv("SENDER_IDENTITY_PARAMETER")
	if emailSenderIdentityParameter == "" {
		panic("SENDER_IDENTITY_PARAMETER env var is not set!")
	}

	logger.WithField("emailSenderIdentityParameter", emailSenderIdentityParameter).Info("Searching email sender identity parameter in SSM")
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
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("Initializing SES client")
	sesClient = ses.NewFromConfig(cfg)

	logger.Info("Initializing SSM client")
	ssmClient = ssm.NewFromConfig(cfg)

	emailSenderIdentity, err := getEmailSenderIdentityParameter(logger)
	if err != nil {
		panic(err)
	}
	logger.Info("Sender email identity", emailSenderIdentity)
}

func main() {
	lambda.Start(handler)
}
