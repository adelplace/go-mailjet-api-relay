package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func sendEndpointRequest(is *is.I, method string, data interface{}) (*response, *httptest.ResponseRecorder) {
	buffer := bytes.NewBufferString("")
	req, err := http.NewRequest(method, "/", buffer)
	is.NoErr(err)

	app := application{}
	w := httptest.NewRecorder()
	app.index(w, req)

	var response = &response{}
	json.NewDecoder(w.Body).Decode(&response)

	return response, w
}

func TestWrongMethod(t *testing.T) {
	is := is.New(t)

	var data struct{}
	response, recorder := sendEndpointRequest(is, http.MethodGet, data)

	is.Equal(recorder.Code, http.StatusMethodNotAllowed)
	is.Equal(response.Code, "method_not_allowed")
	is.Equal(response.Success, false)
}

func TestNoCaptcha(t *testing.T) {
	is := is.New(t)

	var data struct{}
	response, recorder := sendEndpointRequest(is, http.MethodPost, data)

	is.Equal(recorder.Code, http.StatusBadRequest)
	is.Equal(response.Code, "no_captcha")
	is.Equal(response.Success, false)
}
