package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mailjet/mailjet-apiv3-go"
)

type application struct {
	errorLog        *log.Logger
	infoLog         *log.Logger
	httpClient      *http.Client
	mailjetClient   *mailjet.Client
	reCaptchaURL    string
	email           string
	recaptchaSecret string
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network adress")
	recaptchaSecret := os.Getenv("RECAPTCHA_SECRET")
	mailjetPublicKey := os.Getenv("MAILJET_PUBLIC_KEY")
	mailjetPrivateKey := os.Getenv("MAILJET_PRIVATE_KEY")
	email := os.Getenv("EMAIL")
	flag.Parse()

	app := newApplication(recaptchaSecret, mailjetPrivateKey, mailjetPublicKey, email)
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	fmt.Println("Listening")
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}

func newApplication(recaptchaSecret string, mailjetPrivateKey string, mailjetPublicKey string, email string) *application {
	errorLog := log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	httpClient := http.DefaultClient

	return &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		httpClient:      httpClient,
		mailjetClient:   mailjet.NewMailjetClient(mailjetPublicKey, mailjetPrivateKey),
		reCaptchaURL:    "https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s",
		recaptchaSecret: recaptchaSecret,
		email:           email,
	}
}
