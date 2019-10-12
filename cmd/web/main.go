package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	httpClient   *http.Client
	reCaptchaURL string
}

func main() {
	app := newApplication()

	addr := flag.String("addr", ":8080", "HTTP network adress")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	fmt.Println("Listening")
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}

func newApplication() *application {
	errorLog := log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	httpClient := http.DefaultClient

	return &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		httpClient:   httpClient,
		reCaptchaURL: "https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s",
	}
}
