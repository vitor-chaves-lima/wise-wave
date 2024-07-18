package adapters

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"wisewave.tech/email_sender_service/internal/application/usecases"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SQSQueueMessageConsumer struct {
	sendEmailUseCase *usecases.SendEmailUseCase
}

func NewSQSQueueMessageConsumer(sendEmailUseCase *usecases.SendEmailUseCase) ports.QueueMessageConsumer {
	return &SQSQueueMessageConsumer{sendEmailUseCase}
}

func (c *SQSQueueMessageConsumer) Consume(event interface{}) error {
	sqsEvent, ok := event.(events.SQSEvent)
	if !ok {
		return errors.New("event type assertion failed")
	}

	errorsMap := make(map[string]error)

	for _, record := range sqsEvent.Records {
		err := c.sendEmailUseCase.Execute(record.Body)

		if err != nil {
			errorsMap[record.MessageId] = err
		}
	}

	var errorCount = len(errorsMap)
	if errorCount > 0 {
		fmt.Println(errorsMap)
		// TODO: Add messages with errors to DLQ
	}

	return nil
}
