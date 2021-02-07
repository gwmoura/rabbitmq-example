package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"

	"g.com/logger"
	"g.com/mailer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	from, err := mail.ParseAddress("George Moura <gwmoura@gmail.com>")
	to, err := mail.ParseAddress("George Moura <gwmoura@gmail.com>")
	now := time.Now()

	if now.Second()%10 == 1 {
		to, err = mail.ParseAddress("")
	}
	subject := "Teste de email"
	body := "Conteúdo do arquivo"

	if err != nil {
		errorMsg := fmt.Sprintf("Deu pau no email: %s", err)
		logger.PublishLogMessage(errorMsg, "log.*")
		fmt.Fprint(w, errorMsg)
	} else {
		mailer.Send(from, to, subject, body)
		response := fmt.Sprintf("E-mail enviado para %s", to.Name)
		logger.PublishLogMessage(response, "log.*")
		fmt.Fprint(w, response)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/send", sendMail)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
