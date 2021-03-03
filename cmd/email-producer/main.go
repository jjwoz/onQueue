package main

import (
	"github.com/jjwoz/onQueue/internal/util"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func simpleMailDemo() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	util.CheckErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	util.CheckErr(err)
	defer ch.Close()
	_, err = ch.QueueDeclare(
		"email-queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	util.CheckErr(err)
	err = ch.Publish(
		"emails",        // exchange
		"email.created", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("testme@testme.com"),
		})
	util.CheckErr(err)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	simpleMailDemo()
}
