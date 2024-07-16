package parameter_store

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"wisewave.tech/internal/email_sender/models"
)

func getSenderIdentityParameterName() (sender_identity_parameter_name string) {
	sender_identity_parameter_name = os.Getenv("SENDER_IDENTITY_PARAMETER")
	return sender_identity_parameter_name
}

func SenderIdentityFetcher(ssmClient *ssm.Client) (senderIdentity models.SenderIdentity, err error) {
	sender_identity_parameter_name := getSenderIdentityParameterName()
	if sender_identity_parameter_name == "" {
		return models.SenderIdentity{}, fmt.Errorf("coudln't find SENDER_IDENTITY_PARAMETER env var")
	}

	response, err := ssmClient.GetParameter(context.Background(), &ssm.GetParameterInput{
		Name:           aws.String(sender_identity_parameter_name),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return models.SenderIdentity{}, errors.Join(fmt.Errorf("couldn't fetch parameter %s", sender_identity_parameter_name), err)
	}

	senderIdentity, err = models.UnmarshallSenderIdentity(*response.Parameter.Value)
	if err != nil {
		return models.SenderIdentity{}, err
	}

	return senderIdentity, nil
}
