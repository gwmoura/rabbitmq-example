package main

import (
	"log"
	"os"
)

func main() {
	ch := openChannel()

	declareLogsExchange(ch)

	q := declareLogsQueue(ch, "")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		bindLogsQueue(ch, q.Name, s)
	}

	consumeLogMessagesQueue(ch, q.Name, func(body []byte) { log.Printf(" [x] %s", body) })
}
