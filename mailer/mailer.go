package mailer

import (
	"fmt"
	"net/mail"
	"time"

	"g.com/queue"
)

// Send - send mail
func Send(from *mail.Address, to *mail.Address, subject, body string) (bool, error) {
	template := "Date: %s\n"
	template += "From: %s <%s>\n"
	template += "To: %s <%s>\n"
	template += "Subject: %s\n"
	template += "\n\n"
	template += "%s\n"
	template += "\n====\n"
	now := time.Now()

	msg := fmt.Sprintf(template, now, from.Name, from.Address, to.Name, to.Address, subject, body)

	sendMailToQueue(msg)

	return true, nil
}

func sendMailToQueue(msg string) {
	ch := queue.OpenChannel()

	queue.DeclareExchange(ch, "mailer")

	q := queue.DeclareQueue(ch, "emails")

	routingKey := ""

	queue.BindQueue(ch, "mailer", q.Name, routingKey)
	queue.PublishMessage(ch, "mailer", msg, routingKey)
}
