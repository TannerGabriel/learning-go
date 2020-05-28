package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "email", "password", "smtp-client")

	// Email details
	from := "email"
	to := "email"
	subject := "Hello there!"
	body := "Hello there my friend!"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("email-client", auth, "email", []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
