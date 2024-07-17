package adapters

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"wisewave.tech/email_sender_service/internal/domain"
	"wisewave.tech/email_sender_service/internal/ports"
)

type SimpleTemplateManager struct {
	newUserMagicLinkTemplate *template.Template
	magicLinkTemplate        *template.Template
}

func (a *SimpleTemplateManager) formatNewUserMagicLinkEmail(data map[string]interface{}) (subject string, body string, err error) {
	subject = "Seja bem-vindo"

	if _, ok := data["Link"]; !ok {
		return "", "", fmt.Errorf("invalid data format for NewUserMagicLinkEmail template, must have a 'Link' property")
	}

	var bodyBuffer bytes.Buffer
	if err := a.newUserMagicLinkTemplate.Execute(&bodyBuffer, data); err != nil {
		return "", "", err
	}

	body = bodyBuffer.String()

	return subject, body, nil
}

func (a *SimpleTemplateManager) formatMagicLinkEmail(data map[string]interface{}) (subject string, body string, err error) {
	subject = "Seu acesso à experiência"

	if _, ok := data["Link"]; !ok {
		return "", "", fmt.Errorf("invalid data format for MagicLinkEmail template, must have a 'Link' property")
	}

	var bodyBuffer bytes.Buffer
	if err := a.magicLinkTemplate.Execute(&bodyBuffer, data); err != nil {
		return "", "", err
	}

	body = bodyBuffer.String()

	return subject, body, nil
}

func (a *SimpleTemplateManager) FormatEmail(template domain.EmailTemplate, data map[string]interface{}) (subject string, body string, err error) {
	switch template {
	case domain.NewUserMagicLink:
		return a.formatNewUserMagicLinkEmail(data)
	case domain.MagicLink:
		return a.formatMagicLinkEmail(data)
	default:
		return "", "", fmt.Errorf("invalid e-mail template")
	}
}

func loadNewUserMagicLinkTemplate() (templateInstance *template.Template, err error) {
	templateContent := `
    <html lang="pt-br">
    <head>
        <meta charset="UTF-8">
        <link href="https://fonts.googleapis.com/css2?family=Open+Sans&display=swap" rel="stylesheet">
        <style>
            body {
                font-family: 'Open Sans', sans-serif;
            }
        </style>
    </head>
    <body>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
            <tr>
                <td align="center" bgcolor="#ffffff" style="padding: 40px 0 30px 0;">
                    <h1 style="font-size: 24px; margin: 0;">Bem-vindo à experiência WiseWave!</h1>
                </td>
            </tr>
            <tr>
                <td bgcolor="#ffffff" style="padding: 0px 0px 20px 0px;">
                    <div style="padding: 20px 0px 20px 0px;">
                        <p style="font-size: 16px; margin: 0;">Olá,</p>
                        <p style="font-size: 16px; margin: 0;">Obrigado por se cadastrar. Clique no link abaixo para continuar:</p>
                    </div>
                    <p style="font-size: 16px; margin: 0;">
                        <a href="{{.Link}}" style="background-color: #4CAF50; color: white; padding: 12px 20px; text-align: center; text-decoration: none; display: inline-block; border-radius: 4px;">Acessar experiência</a>
                    </p>
                </td>
            </tr>
            <tr>
                <td bgcolor="#f0f0f0" style="padding: 30px 30px 30px 30px;">
                    <p style="font-size: 14px; margin: 0;">Se você não se cadastrou, ignore este email.</p>
                </td>
            </tr>
        </table>
    </body>
    </html>
    `

	templateInstance, err = template.New("NewUserMagicLinkEmail").Parse(templateContent)
	if err != nil {
		return nil, err
	}

	return templateInstance, nil
}

func loadMagicLinkTemplate() (templateInstance *template.Template, err error) {
	templateContent := `
    <html lang="pt-br">
    <head>
        <meta charset="UTF-8">
        <link href="https://fonts.googleapis.com/css2?family=Open+Sans&display=swap" rel="stylesheet">
        <style>
            body {
                font-family: 'Open Sans', sans-serif;
            }
        </style>
    </head>
    <body>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="600">
            <tr>
                <td bgcolor="#ffffff" style="padding: 0px 0px 20px 0px;">
                    <div style="padding: 20px 0px 20px 0px;">
                        <p style="font-size: 16px; margin: 0;">Olá,</p>
                        <p style="font-size: 16px; margin: 0;">Você solicitou acesso à experiência. Clique no link abaixo para continuar:</p>
                    </div>
                    <p style="font-size: 16px; margin: 0;">
                        <a href="{{.Link}}"
                            style="background-color: #4CAF50; color: white; padding: 12px 20px; text-align: center; text-decoration: none; display: inline-block; border-radius: 4px;">Acessar
                            experiência</a>
                    </p>
                </td>
            </tr>
            <tr>
                <td bgcolor="#f0f0f0" style="padding: 30px 30px 30px 30px;">
                    <p style="font-size: 14px; margin: 0;">Se você não solicitou acesso, ignore este email.
                    </p>
                </td>
            </tr>
        </table>
    </body>
    </html>
    `

	templateInstance, err = template.New("MagicLinkEmail").Parse(templateContent)
	if err != nil {
		return nil, err
	}

	return templateInstance, nil
}

func NewSimpleTemplateManager() (templateManager ports.TemplateManager, err error) {
	newUserMagicLinkTemplate, err := loadNewUserMagicLinkTemplate()
	if err != nil {
		return nil, errors.Join(errors.New("couldn't create NewUserMagicLinkEmail template"), err)
	}

	magicLinkTemplate, err := loadMagicLinkTemplate()
	if err != nil {
		return nil, errors.Join(errors.New("couldn't create MagicLinkEmail template"), err)
	}

	return &SimpleTemplateManager{newUserMagicLinkTemplate, magicLinkTemplate}, nil
}
