package email_sender_service

type EmailSenderServiceMessagePublisher interface {
	SendNewUserMagicLinkEmail(userEmail, magicLink string) error
	SendMagicLinkEmail(userEmail, magicLink string) error
}
