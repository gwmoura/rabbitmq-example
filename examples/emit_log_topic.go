package main

import (
	"log"
	"os"
)

func main() {
	ch := openChannel()

	declareLogsExchange(ch)

	body := bodyFrom(os.Args)
	publishLogMessage(ch, body, severityFrom(os.Args))

	log.Printf(" [x] Sent %s", body)
}
