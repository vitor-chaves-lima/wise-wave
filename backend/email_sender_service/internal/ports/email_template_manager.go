package ports

import "wisewave.tech/email_sender_service/internal/domain"

type TemplateManager interface {
	FormatEmail(template domain.EmailTemplate, data map[string]interface{}) (subject string, body string, err error)
}
