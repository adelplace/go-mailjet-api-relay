package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type contact struct {
	email   string
	name    string
	subject string
	message string
}

type captcha struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

const captchaSecret = "secret"

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.renderError(w, "Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.logError(err)
		app.renderError(w, "Unable to parse form", "invalid_data", http.StatusBadRequest)
		return
	}

	captchaResponse := r.PostForm.Get("g-recaptcha-response")
	if len(captchaResponse) == 0 {
		app.renderError(w, "No captcha sent", "no_captcha", http.StatusBadRequest)
		return
	}

	googleResponse := app.checkCaptcha(captchaResponse)
	if !googleResponse {
		app.renderError(w, "Captcha is invalid", "invalid_captcha", http.StatusBadRequest)
		return
	}

	contact := &contact{
		email:   r.PostForm.Get("email"),
		name:    r.PostForm.Get("name"),
		subject: r.PostForm.Get("subject"),
		message: r.PostForm.Get("message"),
	}
	app.sendMail(contact)

	app.renderSuccess(w, "The email has been sent")
}

func (app *application) checkCaptcha(captchaResponse string) (response bool) {
	captcha := captcha{
		Secret:   captchaSecret,
		Response: captchaResponse,
	}
	resp, err := app.httpClient.Get(fmt.Sprintf(app.reCaptchaURL, captcha.Secret, captcha.Response))
	if err != nil {
		app.logError(err)
		return false
	}

	var googleResponse struct {
		Success bool `json:"success"`
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&googleResponse)

	return googleResponse.Success
}
