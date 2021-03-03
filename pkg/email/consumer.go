package email

import (
	"github.com/jjwoz/onQueue/internal/util"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func EmailSendWaiter() {
	queueName := "email-queue"

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	util.CheckErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	util.CheckErr(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	util.CheckErr(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.CheckErr(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			emailAddress := string(d.Body)
			RegistrationWelcomeSend(emailAddress)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
