package queue

import (
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var err error

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Connect() *amqp.Connection {
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	return conn
}

func OpenChannel() *amqp.Channel {
	conn = Connect()
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	// defer ch.Close()

	return ch
}

func DeclareExchange(ch *amqp.Channel, exchangeName string) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare an exchange")
}

func DeclareQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return q
}

func BindQueue(ch *amqp.Channel, exchangeName string, queueName string, routingKey string) {
	err := ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil)
	FailOnError(err, "Failed to bind a queue")
}

func PublishMessage(ch *amqp.Channel, exchangeName, message string, routingKey string) {
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	FailOnError(err, "Failed to publish a message")
}

func ConsumeMessagesQueue(ch *amqp.Channel, queueName string, callbak func([]byte)) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			callbak(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
