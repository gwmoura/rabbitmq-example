package utils

import (
	"log"
	"os"
)

// AppendToFile add a content to and of file
func AppendToFile(body []byte, filePath string) (*os.File, error) {
	log.Printf(" [writing..] %s", body)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(string(body) + "\n"); err != nil {
		log.Println(err)
	}

	return f, err
}
