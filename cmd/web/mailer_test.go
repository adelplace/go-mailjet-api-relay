package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"

	"github.com/mailjet/mailjet-apiv3-go"
)

func mockSendMailFunc(fail bool) func(req *http.Request) (*http.Response, error) {
	var err error
	if fail {
		err = errors.New("Error")
	}
	return func(req *http.Request) (*http.Response, error) {
		responseRecorder := httptest.NewRecorder()
		responseRecorder.Write([]byte(`{"Messages": []}`))
		return responseRecorder.Result(), err
	}
}
func TestSendMail(t *testing.T) {
	is := is.New(t)
	app := newApplication()
	httpClientMocked := mailjet.NewhttpClientMock(true)
	smtpClientMocked := mailjet.NewSMTPClientMock(true)
	httpClientMocked.SendMailV31Func = mockSendMailFunc(false)
	app.mailjetClient = mailjet.NewClient(httpClientMocked, smtpClientMocked, "custom")

	contact := &contact{
		email:   "georges.abitbol@email.com",
		name:    "Georges Abitbol",
		subject: "Important message",
		message: "Hello",
	}
	err := app.sendMail(contact)
	is.NoErr(err)

	httpClientMocked.SendMailV31Func = mockSendMailFunc(true)
	err = app.sendMail(contact)
	is.True(err != nil)
}
