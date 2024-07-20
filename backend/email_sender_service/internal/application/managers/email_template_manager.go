package managers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"

	"github.com/sirupsen/logrus"

	"wisewave.tech/common/lib"
	"wisewave.tech/email_sender_service/internal/domain"
)

type EmailTemplateManager struct {
	ctx                      context.Context
	newUserMagicLinkTemplate *template.Template
	magicLinkTemplate        *template.Template
}

func (a *EmailTemplateManager) formatNewUserMagicLinkEmail(data map[string]interface{}) (subject string, body string, err error) {
	subject = "Seja bem-vindo"

	if _, ok := data["link"]; !ok {
		return "", "", fmt.Errorf("invalid data format for NewUserMagicLinkEmail template, must have a 'link' property")
	}

	var bodyBuffer bytes.Buffer
	if err := a.newUserMagicLinkTemplate.Execute(&bodyBuffer, data); err != nil {
		return "", "", err
	}

	body = bodyBuffer.String()

	return subject, body, nil
}

func (a *EmailTemplateManager) formatMagicLinkEmail(data map[string]interface{}) (subject string, body string, err error) {
	subject = "Seu acesso à experiência"

	if _, ok := data["link"]; !ok {
		return "", "", fmt.Errorf("invalid data format for MagicLinkEmail template, must have a 'link' property")
	}

	var bodyBuffer bytes.Buffer
	if err := a.magicLinkTemplate.Execute(&bodyBuffer, data); err != nil {
		return "", "", err
	}

	body = bodyBuffer.String()

	return subject, body, nil
}

func (a *EmailTemplateManager) FormatEmail(emailTemplateData domain.EmailTemplateData) (subject string, body string, err error) {
	switch emailTemplateData.Type {
	case domain.NewUserMagicLink:
		return a.formatNewUserMagicLinkEmail(emailTemplateData.Data)
	case domain.MagicLink:
		return a.formatMagicLinkEmail(emailTemplateData.Data)
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
                        <a href="{{.link}}" style="background-color: #4CAF50; color: white; padding: 12px 20px; text-align: center; text-decoration: none; display: inline-block; border-radius: 4px;">Acessar experiência</a>
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
                        <a href="{{.link}}"
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

func NewEmailTemplateManager(ctx context.Context) (templateManager *EmailTemplateManager, err error) {
	fields := logrus.Fields{
		"type":          "manager",
		"name":          "email_template_manager",
		"function_name": "NewEmailTemplateManager",
	}
	logger := lib.LoggerFromContext(ctx).WithFields(fields)

	logger.Info("creating EmailTemplateManager")

	logger.Info("loading NewUserMagicLink template")
	newUserMagicLinkTemplate, err := loadNewUserMagicLinkTemplate()
	if err != nil {
		err = errors.Join(errors.New("couldn't create NewUserMagicLinkEmail template"), err)
		logger.Error(err)
		return nil, err
	}

	logger.Info("loading MagicLinkTemplate template")
	magicLinkTemplate, err := loadMagicLinkTemplate()
	if err != nil {
		err = errors.Join(errors.New("couldn't create MagicLinkEmail template"), err)
		logger.Error(err)
		return nil, err
	}

	return &EmailTemplateManager{ctx, newUserMagicLinkTemplate, magicLinkTemplate}, nil
}
