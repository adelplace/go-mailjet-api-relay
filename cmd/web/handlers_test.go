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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	app.index(w, req)

	var response = &response{}
	json.NewDecoder(w.Body).Decode(&response)

	return response, w
}

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

func TestCheckCaptcha(t *testing.T) {
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
