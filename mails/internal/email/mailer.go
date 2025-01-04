package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Mailer struct {
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

func NewMailer(host, port, username, password string) *Mailer {
	return &Mailer{
		smtpHost: host,
		smtpPort: port,
		auth:     smtp.PlainAuth("", username, password, host),
	}
}

func (m *Mailer) SendMail(to, subject, templatePath string, data interface{}) (err error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return
	}

	msg := []byte("Subject: " + subject + "\n\n" + body.String())
	return smtp.SendMail(m.smtpHost+":"+m.smtpPort, m.auth, "no-reply.testing@gmail.com", []string{to}, msg)
}