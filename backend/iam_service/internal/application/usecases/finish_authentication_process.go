package usecases

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/application/dto"
	"wisewave.tech/iam_service/internal/ports"
)

type FinishAuthenticationUseCase struct {
	logger                  *logrus.Entry
	magicLinkChallengeTable ports.MagicLinkChallengeTable
	identityProvider        ports.IdentityProvider
}

func NewFinishAuthenticationUseCase(ctx context.Context, magicLinkChallengeTable ports.MagicLinkChallengeTable, identityProvider ports.IdentityProvider) *FinishAuthenticationUseCase {
	logger := lib.LoggerFromContext(ctx).WithField("type", "usecase")

	return &FinishAuthenticationUseCase{
		logger:                  logger,
		magicLinkChallengeTable: magicLinkChallengeTable,
		identityProvider:        identityProvider,
	}
}

func (uc *FinishAuthenticationUseCase) Execute(ctx context.Context, magicLinkChallenge string) (*dto.UserSessionData, error) {
	logger := uc.logger.WithField("challenge", magicLinkChallenge)

	challenge, sessionToken, userEmail, err := uc.magicLinkChallengeTable.GetChallenge(magicLinkChallenge)
	if err != nil {
		logger.WithError(err).Error("unable to get magic link challenge")
		return nil, err
	}

	if challenge == "" || sessionToken == "" || userEmail == "" {
		err = errors.New("invalid magic link challenge")
		logger.WithError(err).Error("could not find magic link challenge or session token")
		return nil, err
	}

	userSessionData, err := uc.identityProvider.FinishAuthenticationProcess(userEmail, challenge, sessionToken)
	if err != nil {
		logger.WithError(err).Error("unable to finish authentication process")
		return nil, err
	}

	return userSessionData, nil
}
