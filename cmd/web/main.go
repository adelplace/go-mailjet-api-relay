package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	errorLog := log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	addr := flag.String("addr", ":8080", "HTTP network adress")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	fmt.Println("Listening")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
