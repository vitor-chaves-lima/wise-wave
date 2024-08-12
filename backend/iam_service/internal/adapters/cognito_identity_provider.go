package adapters

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/application/dto"
	"wisewave.tech/iam_service/internal/ports"
)

type CognitoIdentityProvider struct {
	logger              *logrus.Entry
	cognitoClient       *cognitoidentityprovider.Client
	userPoolId          string
	applicationClientId string
}

func NewCognitoIdentityProvider(ctx context.Context, cognitoClient *cognitoidentityprovider.Client, userPoolId string, applicationClientId string) ports.IdentityProvider {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "adapter",
		"port": "identity_provider",
	})

	return &CognitoIdentityProvider{
		logger:              logger,
		cognitoClient:       cognitoClient,
		userPoolId:          userPoolId,
		applicationClientId: applicationClientId,
	}
}

func (c *CognitoIdentityProvider) InitiateAuthenticationProcess(userEmail string) (challenge string, authenticationSessionToken string, err error) {
	logger := c.logger.WithField("userEmail", userEmail)

	logger.Info("generating authentication process input")
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: types.AuthFlowTypeCustomAuth,
		AuthParameters: map[string]string{
			"USERNAME": userEmail,
		},
		UserPoolId: &c.userPoolId,
		ClientId:   &c.applicationClientId,
	}

	logger.Info("starting authentication process")
	resp, err := c.cognitoClient.AdminInitiateAuth(context.Background(), params)
	if err != nil {
		c.logger.WithError(err).Error("unable to start authentication process")
		return "", "", err
	}

	challengeName := resp.ChallengeParameters["challenge"]
	if challengeName == "" {
		c.logger.Error("challenge name not found")
		return "", "", errors.New("challenge name not found")
	}

	return challengeName, *resp.Session, nil
}

func (c *CognitoIdentityProvider) CheckUserExists(userEmail string) (bool, error) {
	logger := c.logger.WithField("userEmail", userEmail)

	logger.Info("checking if user exists")
	params := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &c.userPoolId,
		Username:   &userEmail,
	}

	logger.Info("getting user")
	user, err := c.cognitoClient.AdminGetUser(context.Background(), params)
	if err != nil {
		logger.WithError(err).Error("unable to get user")
		var reqErr *types.UserNotFoundException
		if !errors.As(err, &reqErr) {
			return false, err
		} else {
			logger.Info("user does not exist")
			return false, nil
		}
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}

func (c *CognitoIdentityProvider) CheckUserVerified(userId string) (bool, error) {
	logger := c.logger.WithField("userId", userId)

	logger.Info("checking if user is verified")
	params := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &c.userPoolId,
		Username:   &userId,
	}

	logger.Info("getting user")
	user, err := c.cognitoClient.AdminGetUser(context.Background(), params)
	if err != nil {
		logger.WithError(err).Error("unable to get user")
		return false, err
	}

	if user.UserStatus == types.UserStatusTypeConfirmed {
		return true, nil
	}

	return false, nil
}

func (c *CognitoIdentityProvider) AddUser(userEmail string) error {
	logger := c.logger.WithField("userEmail", userEmail)

	logger.Info("adding user")
	params := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:    &c.userPoolId,
		Username:      &userEmail,
		MessageAction: types.MessageActionTypeSuppress,
	}

	logger.Info("creating user")
	_, err := c.cognitoClient.AdminCreateUser(context.Background(), params)
	if err != nil {
		logger.WithError(err).Error("unable to create user")
		return err
	}

	return nil
}

func (c *CognitoIdentityProvider) VerifyUser(userId string) (err error) {
	logger := c.logger.WithField("userId", userId)

	logger.Info("verifying user")
	params := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
		UserPoolId: &c.userPoolId,
		Username:   &userId,
	}

	logger.Info("updating user attributes")
	_, err = c.cognitoClient.AdminUpdateUserAttributes(context.Background(), params)
	if err != nil {
		logger.WithError(err).Error("unable to update user attributes")
		return err
	}

	return nil
}

func (c *CognitoIdentityProvider) FinishAuthenticationProcess(userEmail string, challenge string, sessionToken string) (sessionData *dto.UserSessionData, err error) {
	logger := c.logger.WithField("challenge", challenge)

	logger.Info("finishing authentication process")
	params := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		ChallengeName: types.ChallengeNameTypeCustomChallenge,
		ChallengeResponses: map[string]string{
			"USERNAME": userEmail,
			"ANSWER":   challenge,
		},
		Session:    &sessionToken,
		UserPoolId: &c.userPoolId,
		ClientId:   &c.applicationClientId,
	}

	logger.Info("responding to authentication challenge")
	authChallengeResponse, err := c.cognitoClient.AdminRespondToAuthChallenge(context.Background(), params)
	if err != nil {
		logger.WithError(err).Error("unable to respond to authentication challenge")
		return nil, err
	}

	return &dto.UserSessionData{
		IdToken:      *authChallengeResponse.AuthenticationResult.IdToken,
		AccessToken:  *authChallengeResponse.AuthenticationResult.AccessToken,
		RefreshToken: *authChallengeResponse.AuthenticationResult.RefreshToken,
		ExpiresIn:    authChallengeResponse.AuthenticationResult.ExpiresIn,
		TokenType:    *authChallengeResponse.AuthenticationResult.TokenType,
	}, nil
}
