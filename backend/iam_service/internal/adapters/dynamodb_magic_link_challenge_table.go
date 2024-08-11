package adapters

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/ports"
)

type DynamoDBMagicLinkChallangeTable struct {
	logger         *logrus.Entry
	dynamodbClient *dynamodb.Client
	tableName      string
	challengeTTL   int64
}

func NewDynamodbMagicLinkChallangeTable(ctx context.Context, dynamodbClient *dynamodb.Client, challengeTTL int64, tableName string) ports.MagicLinkChallengeTable {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "adapter",
		"port": "magic_link_challenge_table",
	})

	return &DynamoDBMagicLinkChallangeTable{
		logger,
		dynamodbClient,
		tableName,
		challengeTTL,
	}
}

func (a *DynamoDBMagicLinkChallangeTable) StoreChallenge(challenge string) (err error) {
	logger := a.logger

	completeChallengeTTL := time.Now().Unix() + a.challengeTTL

	logger.Info("generating table data input")
	putItemDataInput := &dynamodb.PutItemInput{
		TableName: aws.String(a.tableName),
		Item: map[string]types.AttributeValue{
			"Challenge": &types.AttributeValueMemberS{Value: challenge},
			"TTL":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", completeChallengeTTL)},
		},
	}

	logger.Info("storing magic link challenge")
	_, err = a.dynamodbClient.PutItem(context.Background(), putItemDataInput)
	if err != nil {
		err := errors.Join(errors.New("couldn't add magic link challenge to table"), err)
		logger.Error(err)
		return err
	}

	return nil
}

func (a *DynamoDBMagicLinkChallangeTable) AssignSessionTokenToChallenge(challenge string, sessionToken string, userEmail string) (err error) {
	logger := a.logger

	completeChallengeTTL := time.Now().Unix() + a.challengeTTL

	logger.Info("generating table data input")
	putItemDataInput := &dynamodb.PutItemInput{
		TableName: aws.String(a.tableName),
		Item: map[string]types.AttributeValue{
			"Challenge":    &types.AttributeValueMemberS{Value: challenge},
			"SessionToken": &types.AttributeValueMemberS{Value: sessionToken},
			"UserEmail":    &types.AttributeValueMemberS{Value: userEmail},
			"TTL":          &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", completeChallengeTTL)},
		},
	}

	logger.Info("assigning session to challenge")
	_, err = a.dynamodbClient.PutItem(context.Background(), putItemDataInput)
	if err != nil {
		err := errors.Join(errors.New("couldn't add magic link challenge to table"), err)
		logger.Error(err)
		return err
	}

	return nil
}

func (a *DynamoDBMagicLinkChallangeTable) GetChallenge(challenge string) (storedChallenge string, storedSessionToken string, userEmail string, err error) {
	logger := a.logger

	logger.Info("generating get item input data")
	getItemDataInput := &dynamodb.GetItemInput{
		TableName: aws.String(a.tableName),
		Key: map[string]types.AttributeValue{
			"Challenge": &types.AttributeValueMemberS{Value: challenge},
		},
	}

	logger.Info("fetching magic link challenge")
	result, err := a.dynamodbClient.GetItem(context.Background(), getItemDataInput)
	if err != nil {
		err := errors.Join(errors.New("couldn't fetch magic link challenge from table"), err)
		logger.Error(err)
		return "", "", "", err
	}

	if result.Item == nil {
		return "", "", "", nil
	}

	magicLinkChallengeAttribute, ok := result.Item["Challenge"].(*types.AttributeValueMemberS)
	if !ok {
		err := fmt.Errorf("failed to get Challenge attribute")
		logger.Error(err)
		return "", "", "", err
	}

	sessionTokenAttribute, ok := result.Item["SessionToken"].(*types.AttributeValueMemberS)
	if !ok {
		err := fmt.Errorf("failed to get SessionToken attribute")
		logger.Error(err)
		return "", "", "", err
	}

	userEmailAttribute, ok := result.Item["UserEmail"].(*types.AttributeValueMemberS)
	if !ok {
		err := fmt.Errorf("failed to get UserEmail attribute")
		logger.Error(err)
		return "", "", "", err
	}

	return magicLinkChallengeAttribute.Value, sessionTokenAttribute.Value, userEmailAttribute.Value, nil
}

func (a *DynamoDBMagicLinkChallangeTable) DeleteChallenge(challenge string) (err error) {
	logger := a.logger

	logger.Info("generating delete item input data")
	deleteItemDataInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(a.tableName),
		Key: map[string]types.AttributeValue{
			"Challenge": &types.AttributeValueMemberS{Value: challenge},
		},
	}

	logger.Info("deleting magic link challenge")
	_, err = a.dynamodbClient.DeleteItem(context.Background(), deleteItemDataInput)
	if err != nil {
		err := errors.Join(errors.New("couldn't delete magic link challenge from table"), err)
		logger.Error(err)
		return err
	}

	return nil
}
