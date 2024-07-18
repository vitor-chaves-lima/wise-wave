package dto

var EmailTypeStrings = []string{
	"newUserMagicLink",
	"magicLink",
}

type EmailMessageDTO struct {
	EmailType string                 `json:"emailType"`
	To        string                 `json:"to"`
	Data      map[string]interface{} `json:"data"`
}
