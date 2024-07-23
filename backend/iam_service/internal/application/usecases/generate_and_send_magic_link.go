package usecases

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	email_sender_service "wisewave.tech/email_sender_service/lib/ports"
	"wisewave.tech/iam_service/internal/ports"
)

type GenerateAndSendMagicLinkUseCase struct {
	logger                      *logrus.Entry
	magicLinkTable              ports.MagicLinkChallengeTable
	emailSenderMessagePublisher email_sender_service.EmailSenderServiceMessagePublisher
}

func NewGenerateAndSendMagicLinkUseCase(ctx context.Context, magicLinkTable ports.MagicLinkChallengeTable, emailSenderMessagePublisher email_sender_service.EmailSenderServiceMessagePublisher) *GenerateAndSendMagicLinkUseCase {
	logger := logrus.WithField("type", "usecase")

	return &GenerateAndSendMagicLinkUseCase{
		logger,
		magicLinkTable,
		emailSenderMessagePublisher,
	}
}

func (uc *GenerateAndSendMagicLinkUseCase) Execute(ctx context.Context, userId string, userEmail string) error {
	logger := uc.logger.WithField("userId", userId)

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

	magicLink := fmt.Sprintf("https://iam.wisewave.tech/magic-link?challenge=%s", magicLinkChallenge)

	logger.Info("sending magic link email")
	err = uc.emailSenderMessagePublisher.SendMagicLinkEmail(userEmail, magicLink)
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
