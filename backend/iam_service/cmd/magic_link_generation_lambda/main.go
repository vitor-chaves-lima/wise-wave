package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event *events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
