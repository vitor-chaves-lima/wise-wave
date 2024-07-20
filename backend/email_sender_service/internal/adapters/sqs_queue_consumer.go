package adapters

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/email_sender_service/internal/application/usecases"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SQSQueueMessageConsumer struct {
	ctx              context.Context
	sendEmailUseCase *usecases.SendEmailUseCase
}

func NewSQSQueueMessageConsumer(ctx context.Context, sendEmailUseCase *usecases.SendEmailUseCase) ports.QueueMessageConsumer {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "adapter",
		"port": "message_consumer",
	})
	ctx = lib.WithLogger(ctx, logger)

	return &SQSQueueMessageConsumer{ctx, sendEmailUseCase}
}

func (c *SQSQueueMessageConsumer) Consume(event interface{}) error {
	logger := lib.LoggerFromContext(c.ctx)
	logger.Info("starting message consumption")

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
		logger.Error(errorsMap)
		panic("couldn't process messages")
	}

	return nil
}
