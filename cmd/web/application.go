package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type response struct {
	Success bool   `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (app *application) logError(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
}

func (app *application) render(w http.ResponseWriter, response *response, statusCode int) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func (app *application) renderSuccess(w http.ResponseWriter, message string) {
	response := &response{
		Success: true,
		Message: message,
		Code:    "success",
	}
	app.render(w, response, http.StatusOK)
}

func (app *application) renderError(w http.ResponseWriter, message string, code string, status int) {
	response := &response{
		Success: false,
		Message: message,
		Code:    code,
	}
	app.render(w, response, status)
}
