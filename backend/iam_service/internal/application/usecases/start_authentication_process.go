package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/application/validators"
	"wisewave.tech/iam_service/internal/ports"
)

type StartAuthenticationProcessUseCase struct {
	logger                  *logrus.Entry
	identityProvider        ports.IdentityProvider
	magicLinkChallengeTable ports.MagicLinkChallengeTable
}

func NewStartAuthenticationProcessUseCase(ctx context.Context, identityProvider ports.IdentityProvider, magicLinkChallengeTable ports.MagicLinkChallengeTable) *StartAuthenticationProcessUseCase {
	logger := lib.LoggerFromContext(ctx).WithFields(logrus.Fields{
		"type": "usecase",
	})

	return &StartAuthenticationProcessUseCase{
		logger:                  logger,
		identityProvider:        identityProvider,
		magicLinkChallengeTable: magicLinkChallengeTable,
	}
}

func (uc *StartAuthenticationProcessUseCase) Execute(ctx context.Context, userEmail string) (err error) {
	logger := uc.logger.WithField("userEmail", userEmail)

	logger.Info("validating email")
	isValidEmail := validators.IsValidEmail(userEmail)
	if !isValidEmail {
		err := &validators.InvalidEmailError{Email: userEmail}
		logger.WithError(err)
		return err
	}

	logger.Info("checking if user exists")
	userExists, err := uc.identityProvider.CheckUserExists(userEmail)
	if err != nil {
		logger.WithError(err).Error("unable to check if user exists")
		return err
	}

	var challenge string
	var authenticationSessionToken string

	if userExists {
		logger.Info("user exists, starting authentication process")
		challenge, authenticationSessionToken, err = uc.identityProvider.InitiateAuthenticationProcess(userEmail)
		if err != nil {
			logger.WithError(err).Error("unable to start authentication process")
			return err
		}
	} else {
		logger.Info("user does not exist, creating user")
		err = uc.identityProvider.AddUser(userEmail)
		if err != nil {
			logger.WithError(err).Error("unable to create user")
			return err
		}

		logger.Info("user created, starting authentication process")
		challenge, authenticationSessionToken, err = uc.identityProvider.InitiateAuthenticationProcess(userEmail)
		if err != nil {
			logger.WithError(err).Error("unable to start authentication process")
			return err
		}
	}

	logger.Info("assigning session token to challenge")
	err = uc.magicLinkChallengeTable.AssignSessionTokenToChallenge(challenge, authenticationSessionToken)
	if err != nil {
		logger.WithError(err).Error("unable to assign session token to challenge")
		return err
	}

	return nil
}
