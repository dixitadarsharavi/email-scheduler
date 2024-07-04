package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	// testing if email sender is provided
	emailSenderResult := "user@domain.com"
	emailSenderExpected := "user@domain.com"
	if emailSenderResult != emailSenderExpected {
		t.Error("Please enter valid email", emailSenderExpected, emailSenderResult)
	}

	// testing if app password is provided
	getEnvVars()
	if os.Getenv("GMAIL_APP_PASSWORD_USER1") == "" {
		t.Error("Please enter valid secret")
	}

	// testing if message is provided
	emailMsgResult := "Service XYZ has planned maintenance on Saturday from 2pm till 5pm cet"
	if emailMsgResult == "" {
		t.Error("Please write a message before sending", "Ex:Service XYZ has planned maintenance on Saturday from 2pm till 5pm cet")
	}

	// testing if host-settings is provided
	hostSettingsResult := []string{"hostName", "port"}
	for index, val := range hostSettingsResult {
		if val == "" {
			t.Error("Please enter valid hostSettings: index", index, "is empty")
		}
	}

}
