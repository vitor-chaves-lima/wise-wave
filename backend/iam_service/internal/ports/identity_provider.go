package ports

import "wisewave.tech/iam_service/internal/application/dto"

type IdentityProvider interface {
	CheckUserExists(userEmail string) (exists bool, err error)
	CheckUserVerified(userId string) (verified bool, err error)
	AddUser(userEmail string) (err error)
	VerifyUser(userId string) (err error)
	InitiateAuthenticationProcess(userEmail string) (challenge string, authenticationSessionToken string, err error)
	FinishAuthenticationProcess(userEmail string, challenge string, sessionToken string) (userSessionData *dto.UserSessionData, err error)
	RefreshToken(refreshToken string) (sessionData *dto.UserSessionData, err error)
}
