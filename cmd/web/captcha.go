package main

import (
	"encoding/json"
	"fmt"
)

type captcha struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

const captchaSecret = "secret"

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
