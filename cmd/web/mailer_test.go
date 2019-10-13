package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"

	"github.com/mailjet/mailjet-apiv3-go"
)

func mockSendMailFunc(success bool) func(req *http.Request) (*http.Response, error) {
	var err error
	if !success {
		err = errors.New("Error")
	}
	return func(req *http.Request) (*http.Response, error) {
		responseRecorder := httptest.NewRecorder()
		responseRecorder.Write([]byte(`{"Messages": []}`))
		return responseRecorder.Result(), err
	}
}

func newMailjetClientMock(success bool) *mailjet.Client {
	httpClientMocked := mailjet.NewhttpClientMock(true)
	smtpClientMocked := mailjet.NewSMTPClientMock(true)
	httpClientMocked.SendMailV31Func = mockSendMailFunc(success)
	return mailjet.NewClient(httpClientMocked, smtpClientMocked, "custom")
}

func TestSendMail(t *testing.T) {
	is := is.New(t)
	app := newApplication()
	contact := &contact{
		email:   "georges.abitbol@email.com",
		name:    "Georges Abitbol",
		subject: "Important message",
		message: "Hello",
	}

	app.mailjetClient = newMailjetClientMock(true)
	err := app.sendMail(contact)
	is.NoErr(err)

	app.mailjetClient = newMailjetClientMock(false)
	err = app.sendMail(contact)
	is.True(err != nil)
}
