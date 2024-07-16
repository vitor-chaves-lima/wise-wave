package ports

type Emailer interface {
	SendEmail(to, subject, body string) error
}
