package mail

import (
	"log"
	"net/smtp"
)

// Mailer is an interface for sending emails
type Mailer interface {
	SendMail(to, content string) error
}

// DevMailer used in dev
type DevMailer struct{}

func (DevMailer) SendMail(to, content string) error {
	log.Printf("Sent email from to %s \n", to)
	log.Printf("Content: %s", content)
	return nil
}

// ProdMailer is the mailer used in production
type ProdMailer struct {
	Host string
	Port string
	User string
	Pass string
	From string
}

func (m *ProdMailer) SendMail(to, content string) error {
	return smtp.SendMail(
		m.Host+":"+m.Port,
		smtp.PlainAuth("", m.User, m.Pass, m.Host),
		m.From,
		[]string{to},
		[]byte(content),
	)

}
