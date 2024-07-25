package ports

type MagicLinkChallengeTable interface {
	StoreChallenge(magicLinkChallenge string) error
	AssignSessionTokenToChallenge(magicLinkChallenge string, sessionToken string) error
	GetChallenge(challenge string) (magicLinkChallenge string, sessionToken string, err error)
	DeleteChallenge(challenge string) (err error)
}
