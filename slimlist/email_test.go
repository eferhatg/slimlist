package slimlist

import (
	"net/smtp"
	"reflect"
	"testing"
)

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

func mockSend(errToReturn error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailRecorder{addr, a, from, to, msg}
		return errToReturn
	}, r
}

func TestNewEmailSender(t *testing.T) {
	ec := EmailConfig{
		Username: "test@test.com",
	}
	es := NewEmailSender(ec)
	if reflect.TypeOf(es).String() != "*slimlist.emailSender" {
		t.Errorf("NewEmailSender type error %v", reflect.TypeOf(es))
	}

}

func TestEmailSend(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{send: f}
	body := "Hello World"
	err := sender.Send([]string{"me@example.com"}, []byte(body))

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(r.msg) != body {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", body, r.msg)
	}
}

func TestSetGlobalEmailSender(t *testing.T) {

	GlobalEmailSender = nil
	SetGlobalEmailSender()
	if GlobalEmailSender == nil {
		t.Errorf("GlobalEmailSender couldn't be set")
	}
}
