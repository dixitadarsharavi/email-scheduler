package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func sendMailSimple(email_sender string, sender_auth string, email_reciever []string, email_msg string, host_settings []string) {
	auth := smtp.PlainAuth(
		"",
		email_sender,
		sender_auth,
		host_settings[0],
	)
	err := smtp.SendMail(
		host_settings[0]+host_settings[1],
		auth,
		email_sender,
		email_reciever,
		[]byte(email_msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	// General APP password created for authentication
	var sender_auth string = "vayn tkje qryv bzww"
	var email_sender string = os.Args[1]
	var email_reciever []string = strings.Split(os.Args[2], ",")
	var email_msg string = os.Args[3]
	var host_settings []string = strings.Split(os.Args[4], ",")
	sendMailSimple(email_sender, sender_auth, email_reciever, email_msg, host_settings)
}
