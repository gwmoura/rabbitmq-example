package logger

import (
	"log"

	"g.com/queue"
	"g.com/utils"
)

func writer(body []byte) {
	_, err := utils.AppendToFile(body, "./tmp/production.log")

	if err != nil {
		log.Println(err)
	}
}

func StartWriter() {
	log.Println("writing log: ...")

	ch := queue.OpenChannel()

	queue.DeclareExchange(ch, "logs")

	q := queue.DeclareQueue(ch, "logs")

	routingKey := "log.*"

	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_topic", routingKey)
	queue.BindQueue(ch, "logs", q.Name, routingKey)

	queue.ConsumeMessagesQueue(ch, q.Name, writer)
}
