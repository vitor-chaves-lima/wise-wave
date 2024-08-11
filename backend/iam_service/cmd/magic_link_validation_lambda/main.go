package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event *events.APIGatewayProxyRequest) (resposne events.APIGatewayProxyResponse, err error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
	}, nil
}

func main() {
	lambda.Start(handler)
}
