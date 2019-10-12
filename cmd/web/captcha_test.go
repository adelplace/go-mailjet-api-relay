package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/matryer/is"
)

func newCpatchaMockHandler(success bool) func(w http.ResponseWriter, r *http.Request) {
	status := "true"
	if success == false {
		status = "false"
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"success":` + status + `}`))
	}
}

func mockCaptchaRequest(app *application, success bool) *httptest.Server {
	mockHandler := newCpatchaMockHandler(success)
	server := httptest.NewServer(http.HandlerFunc(mockHandler))
	app.httpClient = server.Client()
	app.reCaptchaURL = server.URL + "?a=%s&b=%s"

	return server
}

func TestCaptcha(t *testing.T) {
	is := is.New(t)

	app := newApplication()
	data := &url.Values{}
	data.Add("g-recaptcha-response", "fakeToken")

	serverSuccess := mockCaptchaRequest(app, false)
	defer serverSuccess.Close()

	is.Equal(app.checkCaptcha("fakeCaptcha"), false)

	serverFail := mockCaptchaRequest(app, true)
	defer serverFail.Close()

	is.Equal(app.checkCaptcha("fakeCaptcha"), true)
}
