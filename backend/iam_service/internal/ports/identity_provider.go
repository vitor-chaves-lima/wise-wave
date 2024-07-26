package ports

type IdentityProvider interface {
	CheckUserExists(userEmail string) (exists bool, err error)
	AddUser(userEmail string) (err error)
	InitiateAuthenticationProcess(userEmail string) (challenge string, authenticationSessionToken string, err error)
}
