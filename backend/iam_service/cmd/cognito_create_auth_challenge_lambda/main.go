package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/sirupsen/logrus"
)

var (
	ssmClient      *ssm.Client
	dynamodbClient *dynamodb.Client
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsCreateAuthChallenge) {

}

func init() {
	logger := logrus.New().WithField("type", "lambda.init")
	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("loading AWS SDK config")
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	logger.Info("initializing SSM client")
	ssmClient = ssm.NewFromConfig(cfg)

	logger.Info("initializing dynamodb client")
	dynamodbClient = dynamodb.NewFromConfig(cfg)
}

func main() {
	lambda.Start(handler)
}
