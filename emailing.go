package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-gomail/gomail"
)

func sendEmail(envMap map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	color.Blue("-- Please provide Recipient(s) below (You can separate multiple recipients with a comma (NO SPACES!)) --")
	recipient := []string{}
	for len(recipient) == 0 {
		fmt.Print(": ")
		recipientUserInput, err := reader.ReadString('\n')
		if err != nil {
			handleError("Failed to read recipient\nERROR: " + err.Error())
		}
		recipient = strings.Split(strings.Trim(recipientUserInput, "\n"), ",")
		if len(recipient) == 0 {
			color.Red("Recipient cannot be blank!")
		}
	}
	color.Blue("-- Please provide Subject below --")
	subject := ""
	for subject == "" {
		fmt.Print(": ")
		subjectUserInput, err := reader.ReadString('\n')
		if err != nil {
			handleError("Failed to read subject\nERROR: " + err.Error())
		}
		subject = strings.Trim(subjectUserInput, "\n")
		if subject == "" {
			color.Red("Subject cannot be blank!")
		}
	}
	color.Blue("-- Please provide the E-Mail Body below --")
	body := ""
	for body == "" {
		fmt.Print(": ")
		bodyUserInput, err := reader.ReadString('\n')
		if err != nil {
			handleError("Failed to read body\nERROR: " + err.Error())
		}
		body = strings.Trim(bodyUserInput, "\n")
		if body == "" {
			color.Red("Body cannot be blank!")
		}
	}

	m := gomail.NewMessage()
	senderEMail := envMap["Username"] + "@" + envMap["IServInstanceHost"]
	m.SetHeader("From", senderEMail)
	m.SetHeader("To", recipient...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("hgbp.de", 587, envMap["Username"], envMap["Password"])
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
