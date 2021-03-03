package util

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

// Mail represents a mail send request
type Mail struct {
	From       string
	To         []string
	Subject    string
	BodyType   string
	Body       string
	Attachment []string
}

type Subject struct {
	To string
	Product string
}

func SendMail(mail Mail) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To...)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody(mail.BodyType, mail.Body)
	if len(mail.Attachment) > 0 && mail.Attachment != nil {
		for _, attachment := range mail.Attachment {
			m.Attach(attachment)
		}
	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	CheckErr(err)
	d := gomail.NewDialer(os.Getenv("SMTP_HOSTNAME"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"))

	err = d.DialAndSend(m)
	CheckErr(err)
}
