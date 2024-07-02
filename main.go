package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
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
func main() {

	// Sender data
	var emailSender string = os.Args[1]
	var senderAuth string = "vayn tkje qryv bzww" // General APP password created for authentication

	// Receiver email address
	var emailReciever []string = strings.Split(os.Args[2], ",")

	// Message
	var emailMsg string = os.Args[3]

	// Server configuration
	var hostSettings []string = strings.Split(os.Args[4], ",")

	sendMailSimple(emailSender, senderAuth, emailReciever, emailMsg, hostSettings)
}
