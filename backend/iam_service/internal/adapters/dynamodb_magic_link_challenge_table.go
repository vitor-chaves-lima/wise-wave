package adapters

import (
	"context"
	"errors"
	"fmt"

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
		"port": "dynamodb_magic_link_challenge_table",
	})

	return &DynamoDBMagicLinkChallangeTable{
		logger,
		dynamodbClient,
		tableName,
		challengeTTL,
	}
}

func (a *DynamoDBMagicLinkChallangeTable) StoreChallenge(userId string, magicLinkChallenge string) (err error) {
	logger := a.logger.WithField("userId", userId)

	logger.Info("generating table data input")
	putItemDataInput := &dynamodb.PutItemInput{
		TableName: aws.String(a.tableName),
		Item: map[string]types.AttributeValue{
			"UserID":    &types.AttributeValueMemberS{Value: userId},
			"Challenge": &types.AttributeValueMemberS{Value: magicLinkChallenge},
			"TTL":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", a.challengeTTL)},
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

func (a *DynamoDBMagicLinkChallangeTable) GetChallenge(userId string) (magicLinkChallenge string, err error) {
	logger := a.logger.WithField("userId", userId)

	logger.Info("generating get item input data")
	getItemDataInput := &dynamodb.GetItemInput{
		TableName: aws.String(a.tableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userId},
		},
	}

	logger.Info("fetching magic link challenge")
	result, err := a.dynamodbClient.GetItem(context.Background(), getItemDataInput)
	if err != nil {
		err := errors.Join(errors.New("couldn't fetch magic link challenge from table"), err)
		logger.Error(err)
		return "", err
	}

	if result.Item == nil {
		err := fmt.Errorf("magic link challenge not found")
		logger.Error(err)
		return "", err
	}

	magicLinkChallengeAttribute, ok := result.Item["Challenge"].(*types.AttributeValueMemberS)
	if !ok {
		err := fmt.Errorf("failed to get Challenge attribute")
		logger.Error(err)
		return "", err
	}

	return magicLinkChallengeAttribute.Value, nil
}

func (a *DynamoDBMagicLinkChallangeTable) DeleteChallenge(userId string) (err error) {
	logger := a.logger.WithField("userId", userId)

	logger.Info("generating delete item input data")
	deleteItemDataInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(a.tableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userId},
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
