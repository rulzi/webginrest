package helpers

import (
	"webginrest/server/services"

	"github.com/sirupsen/logrus"
	"github.com/wneessen/go-mail"
)

// SendMail sending email
func SendMail(from string, to string, subject string, message string) {
	m := mail.NewMsg()

	if err := m.From(from); err != nil {
		logrus.Fatalf("failed to set From address: %s", err)
	}

	if err := m.To(to); err != nil {
		logrus.Fatalf("failed to set To address: %s", err)
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, message)

	if err := services.Mail.DialAndSend(m); err != nil {
		logrus.Fatalf("failed to send mail: %s", err)
	}
}
