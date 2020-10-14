package mailer

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"net/smtp"
)

func Send(username, from, to, password string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	receivers := []string{to}
	msg := []byte(fmt.Sprintf("To: %v\r\n", to) +
		"Subject: discount 666 Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, receivers, msg)

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Email send to %v from %v successfully", to, from)
}
