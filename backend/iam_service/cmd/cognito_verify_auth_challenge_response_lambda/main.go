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
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/adapters"
	"wisewave.tech/iam_service/internal/application/usecases"
)

var (
	ssmClient      *ssm.Client
	dynamodbClient *dynamodb.Client
)

const (
	magicLinkTableNameEnvVar             = "MAGIC_LINK_TABLE_NAME"
	magicLinkChallengeTTLParameterEnvVar = "MAGIC_LINK_CHALLENGE_TTL_PARAMETER"
)

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
}

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

func loadParameters(logger *logrus.Entry) (challengeTTL int64, magicLinkTableName string, err error) {
	logger = logger.WithField("type", "lambda.loadParameters")

	logger.Info("loading magic link challenge TTL parameter")
	challengeTTLStr, err := getSSMParameterValue(logger, magicLinkChallengeTTLParameterEnvVar)
	if err != nil {
		return 0, "", fmt.Errorf("failed to get magic link challenge TTL: %w", err)
	}
	challengeTTL, err = strconv.ParseInt(challengeTTLStr, 10, 64)
	if err != nil {
		return 0, "", fmt.Errorf("failed to parse magic link challenge TTL: %w", err)
	}

	logger.Info("loading magic link table name parameter")
	magicLinkTableName, err = getEnvVar(magicLinkTableNameEnvVar)
	if err != nil {
		return 0, "", fmt.Errorf("failed to get magic link table name: %w", err)
	}

	return challengeTTL, magicLinkTableName, nil
}

func handler(ctx context.Context, event *events.CognitoEventUserPoolsVerifyAuthChallenge) (*events.CognitoEventUserPoolsVerifyAuthChallenge, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	contextFields := logrus.Fields{
		"requestId":          lambdaContext.AwsRequestID,
		"invokedFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	logger.Info("loading parameters")
	challengeTTL, magicLinkTableName, err := loadParameters(logger)
	if err != nil {
		return event, err
	}

	var response *events.CognitoEventUserPoolsVerifyAuthChallengeResponse = &event.Response

	userId, ok := event.Request.UserAttributes["sub"]
	if !ok {
		return event, fmt.Errorf("sub attribute not found in event")
	}

	challengeAnswer, ok := event.Request.ChallengeAnswer.(string)
	if !ok {
		return event, fmt.Errorf("failed to convert challengeAnswer to string: %w", err)
	}

	magicLinkChallengeTable := adapters.NewDynamodbMagicLinkChallangeTable(ctx, dynamodbClient, challengeTTL, magicLinkTableName)
	generateAndSendMagicLinkUseCase := usecases.NewValidateMagicLinkChallengeUseCase(ctx, magicLinkChallengeTable)

	isTokenValid, err := generateAndSendMagicLinkUseCase.Execute(ctx, challengeAnswer, userId)
	if err != nil {
		return event, fmt.Errorf("failed to validate magic link challenge: %w", err)
	}

	response.AnswerCorrect = isTokenValid

	return event, nil
}

func main() {
	lambda.Start(handler)
}
