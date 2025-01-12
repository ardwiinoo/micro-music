package email

import (
	"bytes"
	"embed"
	"html/template"
	"net/smtp"
)

//go:embed templates/*
var templates embed.FS

type Mailer struct {
	smtpHost string
	smtpPort string
	auth     smtp.Auth
	sender  string
}

func NewMailer(host, port, username, password, sender string) *Mailer {
	return &Mailer{
		smtpHost: host,
		smtpPort: port,
		auth:     smtp.PlainAuth("", username, password, host),
		sender:   sender,
	}
}

func (m *Mailer) SendMail(to, subject, templateName string, data interface{}) error {

	tmplContent, err := templates.ReadFile("templates/" + templateName)
	if err != nil {
		return err
	}

	
	tmpl, err := template.New(templateName).Parse(string(tmplContent))
	if err != nil {
		return err
	}

	
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	
	msg := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			body.String(),
	)
	
	return smtp.SendMail(m.smtpHost+":"+m.smtpPort, m.auth, m.sender, []string{to}, msg)
}
