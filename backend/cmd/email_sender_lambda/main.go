package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"wisewave.tech/internal/email_sender"
	"wisewave.tech/internal/email_sender/models"
	"wisewave.tech/internal/email_sender/parameter_store"
)

var (
	emailSenderService *email_sender.EmailSenderLambdaService
)

func handler(ctx context.Context, event events.SQSEvent) (err error) {
	for _, message := range event.Records {
		messageBody, err := models.UnmarshallMessageBody(message.Body)

		if err != nil {
			return err
		}

		emailSenderService.Execute(message.MessageId, messageBody)
	}

	return nil
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}

	ssmClient := ssm.NewFromConfig(cfg)

	senderIdentity, err := parameter_store.SenderIdentityFetcher(ssmClient)
	if err != nil {
		panic(senderIdentity)
	}

	emailSenderService = email_sender.NewEmailSenderLambdaService(senderIdentity)
}

func main() {
	lambda.Start(handler)
}
