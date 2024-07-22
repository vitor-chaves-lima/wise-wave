package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsDefineAuthChallenge) (response events.CognitoEventUserPoolsDefineAuthChallengeResponse) {
	if len(event.Request.Session) == 0 {
		response.IssueTokens = false
		response.FailAuthentication = false
		response.ChallengeName = "MAGIC_LINK"
	} else {
		lastSession := event.Request.Session[len(event.Request.Session)-1]
		if lastSession.ChallengeName == "MAGIC_LINK" && lastSession.ChallengeResult {
			response.IssueTokens = true
			response.FailAuthentication = false
		} else {
			response.IssueTokens = false
			response.FailAuthentication = true
		}
	}

	return response
}

func main() {
	lambda.Start(handler)
}
