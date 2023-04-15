package mailer

import (
	"fmt"
	"mailService/config"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendMail(toMail string, subject string, body string, config config.Config) error {
	// Sender Data
	senderMail := config.SENDER_GMAIL
	senderPassword := config.SENDER_PASSWORD

	// Receiver Email
	to := []string{
		toMail,
	}

	// smtp configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body = fmt.Sprintf("<h2>%s</h2>", body)

	request := Mail{
		Sender:  senderMail,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	// message
	message := BuildMessage(request)

	// auth
	auth := smtp.PlainAuth("", senderMail, senderPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, to, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
