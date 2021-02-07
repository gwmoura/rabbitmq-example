package logger

import (
	"g.com/queue"
	"github.com/streadway/amqp"
)

func PublishLogMessage(message string, routingKey string) {
	ch := queue.OpenChannel()

	queue.DeclareExchange(ch, "logs")

	q := queue.DeclareQueue(ch, "logs")

	queue.BindQueue(ch, "logs", q.Name, routingKey)
	err := ch.Publish(
		"logs",     // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	queue.FailOnError(err, "Failed to publish a message")
}
