package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
	email_sender_service "wisewave.tech/email_sender_service/lib/adapters"
	"wisewave.tech/iam_service/internal/adapters"
	"wisewave.tech/iam_service/internal/application/usecases"
)

const (
	MagicLinkTableNameEnvVar             = "MAGIC_LINK_TABLE_NAME"
	EmailSenderQueueNameEnvVar           = "EMAIL_SENDER_QUEUE_NAME"
	frontendUrlParameterEnvVar           = "FRONTEND_URL_PARAMETER"
	magicLinkChallengeTTLParameterEnvVar = "MAGIC_LINK_CHALLENGE_TTL_PARAMETER"
)

var (
	ssmClient      *ssm.Client
	dynamodbClient *dynamodb.Client
	sqsClient      *sqs.Client
)

func getSSMParameterValue(logger *logrus.Entry, paramName string) (string, error) {
	value := os.Getenv(paramName)
	if value == "" {
		return "", errors.New(paramName + " environment variable is not set")
	}

	logger.WithField("parameter", value).Info("searching parameter in SSM")
	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(value),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", errors.Join(errors.New("couldn't find parameter in SSM"), err)
	}

	return *param.Parameter.Value, nil
}

func getEnvVar(varName string) (string, error) {
	value := os.Getenv(varName)
	if value == "" {
		return "", fmt.Errorf("%s environment variable is not set", varName)
	}
	return value, nil
}

func loadParameters(logger *logrus.Entry) (frontendUrlParameter string, challengeTTL int64, magicLinkTableName string, emailSenderQueueName string, err error) {
	logger = logger.WithField("type", "lambda.loadParameters")

	logger.Info("loading frontend url parameter")
	frontendUrlParameter, err = getSSMParameterValue(logger, frontendUrlParameterEnvVar)
	if err != nil {
		return "", 0, "", "", fmt.Errorf("failed to get frontend url parameter: %w", err)
	}

	logger.Info("loading magic link challenge TTL parameter")
	challengeTTLStr, err := getSSMParameterValue(logger, magicLinkChallengeTTLParameterEnvVar)
	if err != nil {
		return "", 0, "", "", fmt.Errorf("failed to get magic link challenge TTL: %w", err)
	}
	challengeTTL, err = strconv.ParseInt(challengeTTLStr, 10, 64)
	if err != nil {
		return "", 0, "", "", fmt.Errorf("failed to parse magic link challenge TTL: %w", err)
	}

	logger.Info("loading magic link table name parameter")
	magicLinkTableName, err = getEnvVar(MagicLinkTableNameEnvVar)
	if err != nil {
		return "", 0, "", "", fmt.Errorf("failed to get magic link table name: %w", err)
	}

	logger.Info("loading email sender queue name parameter")
	emailSenderQueueName, err = getEnvVar(EmailSenderQueueNameEnvVar)
	if err != nil {
		return "", 0, "", "", fmt.Errorf("failed to get email sender queue name: %w", err)
	}

	return frontendUrlParameter, challengeTTL, magicLinkTableName, emailSenderQueueName, nil
}

func handler(ctx context.Context, event *events.CognitoEventUserPoolsCreateAuthChallenge) (*events.CognitoEventUserPoolsCreateAuthChallenge, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	contextFields := logrus.Fields{
		"requestId":          lambdaContext.AwsRequestID,
		"invokedFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	logger.Info("loading parameters")
	frontendUrl, challengeTTL, magicLinkTableName, emailSenderQueueName, err := loadParameters(logger)
	if err != nil {
		return event, err
	}

	var response *events.CognitoEventUserPoolsCreateAuthChallengeResponse = &event.Response

	userId, ok := event.Request.UserAttributes["sub"]
	if !ok {
		return event, fmt.Errorf("sub attribute not found in event")
	}

	email, ok := event.Request.UserAttributes["email"]
	if !ok {
		return event, fmt.Errorf("email attribute not found in event")
	}

	emailVerified := "false"
	if val, ok := event.Request.UserAttributes["email_verified"]; ok {
		emailVerified = val
	}

	emailVerifiedBool, err := strconv.ParseBool(emailVerified)
	if err != nil {
		return event, fmt.Errorf("failed to convert emailVerified to boolean: %w", err)
	}

	if len(event.Request.Session) == 0 {
		email_sender_message_publisher := email_sender_service.NewSQSEmailSenderServiceMessagePublisher(ctx, sqsClient, emailSenderQueueName)

		magicLinkChallengeTable := adapters.NewDynamodbMagicLinkChallangeTable(ctx, dynamodbClient, challengeTTL, magicLinkTableName)
		generateAndSendMagicLinkUseCase := usecases.NewGenerateAndSendMagicLinkUseCase(ctx, magicLinkChallengeTable, email_sender_message_publisher, frontendUrl)

		challenge, err := generateAndSendMagicLinkUseCase.Execute(ctx, userId, email, emailVerifiedBool)
		if err != nil {
			return event, fmt.Errorf("failed to generate and send magic link: %w", err)
		}

		response.PublicChallengeParameters = map[string]string{
			"email":     email,
			"challenge": challenge,
		}
	}

	return event, nil
}

func init() {
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("initializing ssm client")
	ssmClient = ssm.NewFromConfig(cfg)

	logger.Info("initializing dynamodb client")
	dynamodbClient = dynamodb.NewFromConfig(cfg)

	logger.Info("initializing sqs client")
	sqsClient = sqs.NewFromConfig(cfg)
}

func main() {
	lambda.Start(handler)
}
