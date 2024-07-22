package ports

type MagicLinkChallengeTable interface {
	StoreChallenge(userId string, magicLinkChallenge string) error
	GetChallenge(userId string) (magicLinkChallenge string, err error)
	DeleteChallenge(userId string) (err error)
}
