package usecases

import (
	"context"

	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	"wisewave.tech/iam_service/internal/application/dto"
	"wisewave.tech/iam_service/internal/ports"
)

type RefreshTokenUseCase struct {
	logger           *logrus.Entry
	identityProvider ports.IdentityProvider
}

func NewRefreshTokenUseCase(ctx context.Context, identityProvider ports.IdentityProvider) *RefreshTokenUseCase {
	logger := lib.LoggerFromContext(ctx).WithField("type", "usecase")

	return &RefreshTokenUseCase{
		logger,
		identityProvider,
	}
}

func (uc *RefreshTokenUseCase) Execute(ctx context.Context, refreshToken string) (userSessionData *dto.UserSessionData, err error) {
	logger := uc.logger

	logger.Info("refreshing token")
	newTokenData, err := uc.identityProvider.RefreshToken(refreshToken)
	if err != nil {
		logger.WithError(err).Error("couldn't refresh token")
		return nil, err
	}

	logger.Info("token refreshed")
	return newTokenData, nil
}
