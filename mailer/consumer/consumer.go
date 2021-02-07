package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"

	"g.com/logger"
	"g.com/queue"
	"g.com/utils"
)

func print(body []byte) {
	log.Printf(" [x] %s", body)

	r := strings.NewReader(string(body))
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	b, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)

	mailPath := fmt.Sprintf("./data/mail/inbox/%s.mail", header.Get("Subject"))
	logger.PublishLogMessage("E-mail enviado", "log.*")
	_, _ = utils.AppendToFile(body, mailPath)
}

func main() {
	fmt.Println("Starting consumer")
	ch := queue.OpenChannel()
	queue.DeclareExchange(ch, "mailer")
	q := queue.DeclareQueue(ch, "emails")
	queue.BindQueue(ch, "mailer", q.Name, "")
	queue.ConsumeMessagesQueue(ch, q.Name, print)
}
