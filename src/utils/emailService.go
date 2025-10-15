package utils

import (
	"html/template"
	"natasha/src/config"

	"github.com/wneessen/go-mail"
)

type EmailData struct {
	ToEmail      string
	Subject      string
	TemplatePath string
	TemplateData map[string]string
}

func SendingEmail(data *EmailData) error {
	// create new a message
	m := mail.NewMsg()
	if err := m.From(config.Data.Email.Address); err != nil {
		return err
	}

	if err := m.To(data.ToEmail); err != nil {
		return err
	}

	m.Subject(data.Subject)

	// load template file from path
	tpl, err := template.ParseFiles(data.TemplatePath)
	if err != nil {
		return err
	}

	// parse data to template
	if err := m.SetBodyHTMLTemplate(tpl, data.TemplateData); err != nil {
		return err
	}

	// new a client
	c, err := mail.NewClient("smtp.gmail.com",
		mail.WithPort(587),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithUsername(config.Data.Email.Address),
		mail.WithPassword(config.Data.Email.Password),
	)

	// send email with client
	return c.DialAndSend(m)

}
