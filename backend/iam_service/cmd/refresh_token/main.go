package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/adapters"
	"wisewave.tech/iam_service/internal/application/usecases"
)

var (
	cognitoClient *cognitoidentityprovider.Client
)

const (
	cognitoUserPoolIdEnvVar          = "COGNITO_USER_POOL_ID"
	cognitoApplicationClientIdEnvVar = "COGNITO_APPLICATION_CLIENT_ID"
)

type RequestBody struct {
	RefreshToken string `json:"refreshToken"`
}

func init() {
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("initializing cognito client")
	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
}

func getEnvVar(varName string) (string, error) {
	value := os.Getenv(varName)
	if value == "" {
		return "", fmt.Errorf("%s environment variable is not set", varName)
	}
	return value, nil
}

func loadParameters() (userPoolId string, applicationClientId string, err error) {
	userPoolId, err = getEnvVar(cognitoUserPoolIdEnvVar)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user pool id: %w", err)
	}

	applicationClientId, err = getEnvVar(cognitoApplicationClientIdEnvVar)
	if err != nil {
		return "", "", fmt.Errorf("failed to get application client id: %w", err)
	}

	return userPoolId, applicationClientId, nil
}

func handler(ctx context.Context, event *events.APIGatewayProxyRequest) (resposne events.APIGatewayProxyResponse, err error) {
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
		log.Printf("failed to unmarshal request body: %v", err)

		responseBody, _ := json.Marshal(map[string]string{
			"message": "Invalid request body",
		})

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date",
				"Access-Control-Allow-Methods": "POST",
			},
			Body: string(responseBody),
		}, nil
	}

	logger.Info("loading parameters")
	userPoolId, applicationClientId, err := loadParameters()
	if err != nil {
		responseBody, _ := json.Marshal(map[string]string{
			"message": "Internal error",
		})

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date",
				"Access-Control-Allow-Methods": "POST",
			},
			Body: string(responseBody),
		}, err
	}

	identityProvider := adapters.NewCognitoIdentityProvider(ctx, cognitoClient, userPoolId, applicationClientId)

	refreshTokenUseCase := usecases.NewRefreshTokenUseCase(ctx, identityProvider)
	userSessionData, err := refreshTokenUseCase.Execute(ctx, body.RefreshToken)
	if err != nil {
		responseBody, _ := json.Marshal(map[string]string{
			"message": "Internal error",
		})

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date",
				"Access-Control-Allow-Methods": "POST",
			},
			Body: string(responseBody),
		}, err
	}

	bodyResponse, err := json.Marshal(userSessionData)
	if err != nil {
		responseBody, _ := json.Marshal(map[string]string{
			"message": "Internal error",
		})

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date",
				"Access-Control-Allow-Methods": "POST",
			},
			Body: string(responseBody),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date",
			"Access-Control-Allow-Methods": "POST",
		},
		Body: string(bodyResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}
