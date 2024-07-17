package domain

type EmailTemplate int

const (
	NewUserMagicLink EmailTemplate = iota
	MagicLink
)
