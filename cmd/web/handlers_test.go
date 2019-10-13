package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func sendEndpointRequest(app *application, is *is.I, method string, data *url.Values) (*response, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	app.index(w, req)

	var response = &response{}
	json.NewDecoder(w.Body).Decode(&response)

	return response, w
}

func TestWrongMethod(t *testing.T) {
	is := is.New(t)

	data := &url.Values{}
	app := &application{}
	response, recorder := sendEndpointRequest(app, is, http.MethodGet, data)

	is.Equal(recorder.Code, http.StatusMethodNotAllowed)
	is.Equal(response.Code, "method_not_allowed")
	is.Equal(response.Success, false)
}

func TestNoCaptcha(t *testing.T) {
	is := is.New(t)

	data := &url.Values{}
	app := &application{}
	response, recorder := sendEndpointRequest(app, is, http.MethodPost, data)

	is.Equal(recorder.Code, http.StatusBadRequest)
	is.Equal(response.Code, "no_captcha")
	is.Equal(response.Success, false)
}

func TestCompleteSend(t *testing.T) {
	is := is.New(t)

	data := &url.Values{}
	data.Add("g-recaptcha-response", "fakeToken")
	data.Add("name", "Georges Abitbol")
	data.Add("email", "georges.abitbol@email.com")
	data.Add("subject", "Very important message")
	data.Add("message", "Hello")

	app := newApplication()

	serverSuccess := mockCaptchaRequest(app, true)
	defer serverSuccess.Close()

	app.mailjetClient = newMailjetClientMock(true)
	response, recorder := sendEndpointRequest(app, is, http.MethodPost, data)

	is.Equal(response.Code, "success")
	is.Equal(recorder.Code, http.StatusOK)
	is.Equal(response.Success, true)

	app.mailjetClient = newMailjetClientMock(false)
	response, recorder = sendEndpointRequest(app, is, http.MethodPost, data)

	is.Equal(response.Code, "mailjet_error")
	is.Equal(recorder.Code, http.StatusBadRequest)
	is.Equal(response.Success, false)
}
