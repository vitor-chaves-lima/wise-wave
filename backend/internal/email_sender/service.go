package email_sender

import (
	"fmt"

	"wisewave.tech/internal/email_sender/models"
)

type EmailSenderLambdaService struct {
	SenderIdentity models.SenderIdentity
}

func (s *EmailSenderLambdaService) Execute(messageId string, messageBody models.MessageBodyBase) {
	fmt.Printf("%+v\n", s.SenderIdentity)
	fmt.Printf("Message ID: %s -> %+v", messageId, messageBody)
}

func NewEmailSenderLambdaService(senderIdentity models.SenderIdentity) *EmailSenderLambdaService {
	return &EmailSenderLambdaService{
		SenderIdentity: senderIdentity,
	}
}
