package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func sendMailSimple(emailSender string, senderAuth string, emailReciever []string, emailMsg string, hostSettings []string) {
	auth := smtp.PlainAuth(
		"",
		emailSender,
		senderAuth,
		hostSettings[0],
	)
	err := smtp.SendMail(
		hostSettings[0]+hostSettings[1],
		auth,
		emailSender,
		emailReciever,
		[]byte(emailMsg),
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully!")
}

func getEnvVars() {
	// load .env file
	err := godotenv.Load("credentials.env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// Sender data
	var emailSender string = os.Args[1]
	getEnvVars()
	var senderAuth string = os.Getenv("GMAIL_APP_PASSWORD_USER1") // General APP password created for authentication
	fmt.Println(senderAuth)

	// Receiver email address
	var emailReciever []string = strings.Split(os.Args[2], ",")

	// Message
	var emailMsg string = os.Args[3]

	// Server configuration
	var hostSettings []string = strings.Split(os.Args[4], ",")

	sendMailSimple(emailSender, senderAuth, emailReciever, emailMsg, hostSettings)
}
