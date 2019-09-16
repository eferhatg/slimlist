package slimlist

import (
	"errors"
	"net/smtp"
	"os"
	"sync"
)

//GlobalEmailSender is the
var (
	GlobalEmailSender *EmailSender
	once              sync.Once
)

//SetGlobalEmailSender sets the global emailsender object
func SetGlobalEmailSender() (*EmailSender, error) {

	if len(os.Getenv("SMTP_USERNAME")) == 0 ||
		len(os.Getenv("SMTP_PASSWORD")) == 0 ||
		len(os.Getenv("SMTP_HOST")) == 0 ||
		len(os.Getenv("SMTP_PORT")) == 0 ||
		len(os.Getenv("SMTP_SENDER")) == 0 {
		return nil, errors.New("SMTP environment variables couldn't be found")
	}

	once.Do(func() {
		ec := EmailConfig{
			Username:   os.Getenv("SMTP_USERNAME"),
			Password:   os.Getenv("SMTP_PASSWORD"),
			ServerHost: os.Getenv("SMTP_HOST"),
			ServerPort: os.Getenv("SMTP_PORT"),
			SenderAddr: os.Getenv("SMTP_SENDER"),
		}
		es := NewEmailSender(ec)

		GlobalEmailSender = &es
	})
	return GlobalEmailSender, nil
}

//EmailConfig keeps email configuration
type EmailConfig struct {
	Username   string
	Password   string
	ServerHost string
	ServerPort string
	SenderAddr string
}

//EmailSender is interface for emailSender
type EmailSender interface {
	Send(to []string, body []byte) error
}

//NewEmailSender inits emailSender
func NewEmailSender(conf EmailConfig) EmailSender {
	return &emailSender{conf, smtp.SendMail}
}

type emailSender struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}

//Send sends email
func (e *emailSender) Send(to []string, body []byte) error {
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, to, body)
}
