package domain

type EmailType int

const (
	NewUserMagicLink EmailType = iota
	MagicLink
)

type EmailTemplateData struct {
	Type EmailType
	Data map[string]interface{}
}
