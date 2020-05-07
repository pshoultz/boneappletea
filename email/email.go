package email

import (
	"log"
	"net/smtp"
)

func Send(address string, message string) {
	log.Println(address, message)
	from := "use bat email here"
	pass := "use bat email pw here"
	to := address

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
