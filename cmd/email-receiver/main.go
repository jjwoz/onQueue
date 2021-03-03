package main

import (
	"github.com/jjwoz/onQueue/pkg/email"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	email.EmailSendWaiter()
}
