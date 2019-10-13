package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"

	"github.com/mailjet/mailjet-apiv3-go"
)

func TestSendMail(t *testing.T) {
	is := is.New(t)
	app := newApplication()
	httpClientMocked := mailjet.NewhttpClientMock(true)
	smtpClientMocked := mailjet.NewSMTPClientMock(true)
	httpClientMocked.SendMailV31Func = func(req *http.Request) (*http.Response, error) {
		responseRecorder := httptest.NewRecorder()
		responseRecorder.Write([]byte(`{"Messages": []}`))
		return responseRecorder.Result(), nil
	}
	app.mailjetClient = mailjet.NewClient(httpClientMocked, smtpClientMocked, "custom")

	contact := &contact{
		email:   "email@email.com",
		name:    "Georges Abitbol",
		subject: "L'homme le plus classe du monde",
		message: "lorem",
	}
	err := app.sendMail(contact)
	is.NoErr(err)
}
