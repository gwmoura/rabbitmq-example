package logger

import (
	"log"

	"g.com/queue"
)

func print(body []byte) {
	log.Printf(" [x] %s", body)
}

func StartConsole() {
	log.Println("printing logs: ...")

	ch := queue.OpenChannel()

	queue.DeclareExchange(ch, "logs")

	q := queue.DeclareQueue(ch, "")

	routingKey := "log.*"

	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_topic", routingKey)

	queue.BindQueue(ch, "logs", q.Name, routingKey)

	queue.ConsumeMessagesQueue(ch, q.Name, print)
}
