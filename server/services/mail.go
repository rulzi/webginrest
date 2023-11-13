package services

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wneessen/go-mail"
)

// Mail variable mail
var Mail *mail.Client

// OpenEmailClient Open Mail Client
func OpenEmailClient(ctx context.Context, mailHost string, mailPort int, mailUser string, mailPass string) {
	c, err := mail.NewClient(mailHost, mail.WithPort(mailPort), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(mailUser), mail.WithPassword(mailPass))
	if err != nil {
		logrus.Warn(err)
	}

	Mail = c
}
