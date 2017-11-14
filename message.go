package main

import (
	"regexp"
	"strings"
)

type Message struct {
	Username         string
	Password         string
	Confirm_Password string
	Email            string
	Name             string
	Telephone        string
	Country          string
	City             string
	Address          string
	Errors           map[string]string
	Result           string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(msg.Email))
	if matched == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please write a valid user name "
	}

	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please write the password "
	} else {
		if strings.TrimSpace(msg.Password) == strings.TrimSpace(msg.Confirm_Password) {
			msg.Errors["Password"] = "Password's don't match "
		}
	}
	if strings.TrimSpace(msg.Telephone) == "" {
		msg.Errors["Telephone"] = "Please write the telephone"
	}

	return len(msg.Errors) == 0
}
