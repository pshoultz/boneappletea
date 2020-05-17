package email

import (
	"log"
	"net/smtp"
)

func Send(address string, message string) {
	//NOTE: i know this is sort of fucked up looking but it works for now.  will fix if it gets popular

	from := "boneappleteam@gmail.com" //NOTE: who sends the email
	pass := "rubberducky"
	to := from //NOTE: who it came from

	msg := "From: " + from + "\n" +
		"To: " + address + "\n" +
		"Subject: Hello there\n\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
