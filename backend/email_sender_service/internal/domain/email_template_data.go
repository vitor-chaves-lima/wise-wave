package domain

type EmailType int

const (
	NewUserMagicLink EmailType = iota
	MagicLink
)

func (e EmailType) String() string {
	return [...]string{"NewUserMagicLink", "MagicLink"}[e]
}

type EmailTemplateData struct {
	Type EmailType
	Data map[string]interface{}
}
