package email

import (
	"example/vtr-mailer-service/actions/email/providers/mailgun"
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
)

func SendEmail(app *application.Application, transaction structs.Transaction, htmlBody string) error {
	var err error
	domain, username, password := app.Cfg.GetMailgunConfig()

	var EmailData = structs.EmailData{
		To:       transaction.To,
		Cc:       transaction.Cc,
		Bcc:      transaction.Bcc,
		From:     transaction.From,
		Subject:  transaction.Subject,
		HtmlBody: htmlBody,
		SendMode: app.Cfg.SendMode,
	}

	// TODO: Multiple email providers
	var mailgunConfig = mailgun.ServiceConfig{
		Domain:   domain,
		Username: username,
		Password: password,
	}

	mailgunResponse, err := mailgun.Dispatch(mailgunConfig, EmailData, htmlBody)

	tools.PrintStruct(mailgunResponse)

	return err
}
