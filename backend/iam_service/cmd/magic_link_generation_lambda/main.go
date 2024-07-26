package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/adapters"
	"wisewave.tech/iam_service/internal/application/usecases"
	"wisewave.tech/iam_service/internal/application/validators"
)

var (
	ssmClient      *ssm.Client
	dynamodbClient *dynamodb.Client
	cognitoClient  *cognitoidentityprovider.Client
)

const (
	magicLinkTableNameEnvVar             = "MAGIC_LINK_TABLE_NAME"
	magicLinkChallengeTTLParameterEnvVar = "MAGIC_LINK_CHALLENGE_TTL_PARAMETER"
	cognitoUserPoolIdEnvVar              = "COGNITO_USER_POOL_ID"
	cognitoApplicationClientIdEnvVar     = "COGNITO_APPLICATION_CLIENT_ID"
)

type RequestBody struct {
	UserEmail string `json:"email"`
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

	logger.Info("initializing cognito client")
	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
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

func loadParameters(logger *logrus.Entry) (challengeTTL int64, magicLinkTableName string, userPoolId string, applicationClientId string, err error) {
	logger = logger.WithField("type", "lambda.loadParameters")

	logger.Info("loading magic link challenge TTL parameter")
	challengeTTLStr, err := getSSMParameterValue(logger, magicLinkChallengeTTLParameterEnvVar)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("failed to get magic link challenge TTL: %w", err)
	}
	challengeTTL, err = strconv.ParseInt(challengeTTLStr, 10, 64)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("failed to parse magic link challenge TTL: %w", err)
	}

	logger.Info("loading magic link table name parameter")
	magicLinkTableName, err = getEnvVar(magicLinkTableNameEnvVar)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("failed to get magic link table name: %w", err)
	}

	userPoolId, err = getEnvVar(cognitoUserPoolIdEnvVar)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("failed to get user pool id: %w", err)
	}

	applicationClientId, err = getEnvVar(cognitoApplicationClientIdEnvVar)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("failed to get application client id: %w", err)
	}

	return challengeTTL, magicLinkTableName, userPoolId, applicationClientId, nil
}

func handler(ctx context.Context, event *events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)

	contextFields := logrus.Fields{
		"requestId":          lambdaContext.AwsRequestID,
		"invokedFunctionArn": lambdaContext.InvokedFunctionArn,
	}

	logger := lib.NewLogger(lib.JSONFormatter).WithField("type", "lambda.handler").WithField("record", contextFields)
	ctx = lib.WithLogger(ctx, logger)

	logger.Info("unmarshalling request body")
	var body RequestBody
	if err := json.Unmarshal([]byte(event.Body), &body); err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	logger.Info("loading parameters")
	challengeTTL, magicLinkTableName, userPoolId, applicationClientId, err := loadParameters(logger)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "{\"message\": \"Internal error\"}",
			StatusCode: 500,
		}, err
	}

	magicLinkChallengeTable := adapters.NewDynamodbMagicLinkChallangeTable(ctx, dynamodbClient, challengeTTL, magicLinkTableName)
	identityProvider := adapters.NewCognitoIdentityProvider(ctx, cognitoClient, userPoolId, applicationClientId)

	startAuthenticationUseCase := usecases.NewStartAuthenticationProcessUseCase(ctx, identityProvider, magicLinkChallengeTable)
	err = startAuthenticationUseCase.Execute(ctx, body.UserEmail)
	if err != nil {
		var reqErr *validators.InvalidEmailError
		if !errors.As(err, &reqErr) {
			logger.WithError(err).Info("internal error")
			return events.APIGatewayProxyResponse{
				Body:       "{\"message\": \"Internal error\"}",
				StatusCode: 500,
			}, err
		} else {
			logger.WithError(err).Info("invalid email")
			return events.APIGatewayProxyResponse{
				Body:       "{\"message\": \"Invalid email\"}",
				StatusCode: 400,
			}, nil
		}
	}

	if err != nil {

		logger.WithError(err).Error("unable to start authentication process")
		return events.APIGatewayProxyResponse{
			Body:       "{\"message\": \"Internal error\"}",
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}

func main() {
	lambda.Start(handler)
}
