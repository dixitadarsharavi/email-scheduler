package main

import (
	"fmt"
	"net/smtp"
)

// Parameters:
// - Email recipients (array)
// - Email sender (1 value)
// - Email message
// - Email server configuration(?) -- may be create a object which has server configuration

func sendMailSimple(email_sender string, sender_auth string, email_reciever []string, msg string, host_name string) {
	auth := smtp.PlainAuth(
		"",
		email_sender,
		sender_auth,
		host_name,
	)

	// msg := "Subject: My special subject\nThis is the body of my email"

	err := smtp.SendMail(
		host_name+":587",
		auth,
		email_sender,
		email_reciever,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	sendMailSimple("adarshravidixit@gmail.com", "vayn tkje qryv bzww", []string{"adarshravidixit@gmail.com", "dixitadarsharavi@gmail.com"}, "Subject: My special subject\nThis is the body of my email", "smtp.gmail.com")
}
