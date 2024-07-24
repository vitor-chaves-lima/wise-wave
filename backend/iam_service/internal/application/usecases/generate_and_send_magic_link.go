package usecases

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wisewave.tech/common/lib"
	email_sender_service "wisewave.tech/email_sender_service/lib/ports"
	"wisewave.tech/iam_service/internal/ports"
)

type GenerateAndSendMagicLinkUseCase struct {
	logger                      *logrus.Entry
	magicLinkTable              ports.MagicLinkChallengeTable
	emailSenderMessagePublisher email_sender_service.EmailSenderServiceMessagePublisher
	frontendUrl                 string
}

func NewGenerateAndSendMagicLinkUseCase(ctx context.Context, magicLinkTable ports.MagicLinkChallengeTable, emailSenderMessagePublisher email_sender_service.EmailSenderServiceMessagePublisher, frontendUrl string) *GenerateAndSendMagicLinkUseCase {
	logger := lib.LoggerFromContext(ctx).WithField("type", "usecase")

	return &GenerateAndSendMagicLinkUseCase{
		logger,
		magicLinkTable,
		emailSenderMessagePublisher,
		frontendUrl,
	}
}

func (uc *GenerateAndSendMagicLinkUseCase) Execute(ctx context.Context, userId string, userEmail string, emailVerified bool) error {
	logger := uc.logger.WithField("userId", userId).WithField("userEmail", userEmail)

	logger.Info("generating magic link challenge")
	magicLinkChallenge, err := generateMagicLinkChallenge(userId)
	if err != nil {
		logger.WithError(err).Error("couldn't generate magic link challenge")
		return err
	}

	logger.Info("storing magic link challenge")
	err = uc.magicLinkTable.StoreChallenge(userId, magicLinkChallenge)
	if err != nil {
		logger.WithError(err).Error("couldn't store magic link challenge")
		return err
	}

	magicLink := fmt.Sprintf("%s/magic-link/validate?challenge=%s", uc.frontendUrl, magicLinkChallenge)

	if emailVerified {
		logger.Info("user is verified, sending magic link email")
		err = uc.emailSenderMessagePublisher.SendMagicLinkEmail(userEmail, magicLink)
	} else {
		logger.Info("user is not verified, sending new user magic link email")
		err = uc.emailSenderMessagePublisher.SendNewUserMagicLinkEmail(userEmail, magicLink)
	}

	if err != nil {
		logger.WithError(err).Error("couldn't send magic link email")
		return err
	}

	return nil
}

func generateMagicLinkChallenge(userId string) (magicLinkChallenge string, err error) {
	id := uuid.New().String()

	timestamp := time.Now().Unix()

	combined := fmt.Sprintf("%s-%d-%s", id, timestamp, userId)

	hash := sha256.New()
	hash.Write([]byte(combined))
	hashedBytes := hash.Sum(nil)

	randomString := hex.EncodeToString(hashedBytes)

	return randomString, nil
}
