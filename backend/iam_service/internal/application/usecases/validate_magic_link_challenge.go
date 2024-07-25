package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/ports"
)

type ValidateMagicLinkChallengeUseCase struct {
	logger         *logrus.Entry
	magicLinkTable ports.MagicLinkChallengeTable
}

func NewValidateMagicLinkChallengeUseCase(ctx context.Context, magicLinkTable ports.MagicLinkChallengeTable) *ValidateMagicLinkChallengeUseCase {
	logger := lib.LoggerFromContext(ctx).WithField("type", "usecase")

	return &ValidateMagicLinkChallengeUseCase{
		logger,
		magicLinkTable,
	}
}

func (uc *ValidateMagicLinkChallengeUseCase) Execute(ctx context.Context, challengeAnswer string, userId string) (bool, error) {
	logger := uc.logger.WithField("userId", userId)

	logger.Info("fetching magic link challenge from table")
	storedMagicLinkChallenge, _, err := uc.magicLinkTable.GetChallenge(challengeAnswer)
	if err != nil {
		logger.WithError(err).Error("couldn't fetch magic link challenge from table")
		return false, err
	}

	if storedMagicLinkChallenge == "" {
		logger.Info("magic link challenge not found in table, magic link challenge is not valid")
		return false, nil
	}

	isMagicLinkChallengeValid := storedMagicLinkChallenge == challengeAnswer
	if isMagicLinkChallengeValid {
		logger.Info("magic link challenge is valid")
	} else {
		logger.Info("magic link challenge is not valid")
	}

	logger.Info("deleting magic link challenge from table")
	err = uc.magicLinkTable.DeleteChallenge(challengeAnswer)
	if err != nil {
		logger.WithError(err).Error("couldn't delete magic link challenge from table")
		return false, err
	}

	return isMagicLinkChallengeValid, nil
}
