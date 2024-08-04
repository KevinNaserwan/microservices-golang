package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const webPort = "80"

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createdMail(),
	}

	log.Println("Starting mail service on port:" + webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func createdMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:     os.Getenv("MAIL_DOMAIN"),
		Host:       os.Getenv("MAIL_HOST"),
		Port:       port,
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
		Encryption: os.Getenv("MAIL_ENCRYPTION"),
		FromName:   os.Getenv("MAIL_FROM_NAME"),
		FromAddr:   os.Getenv("MAIL_FROM_ADDRESS"),
	}

	return m
}
