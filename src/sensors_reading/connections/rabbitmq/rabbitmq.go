package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connect(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)

	if err != nil {
		log.Fatalf("Fail to connect to rabbitmq due to: %s\n", err.Error())
	}
	return conn
}

func openChannel(rabbitUrl string) *amqp.Channel {
	conn := connect(rabbitUrl)

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open the channel due to: %s\n", err.Error())
	}
	return ch
}

func GenerateConsumer(queue string) <-chan amqp.Delivery {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	ch := openChannel(rabbitmqURL)

	log.Println("Channel gerado")

	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	log.Println("Gerando consumidor do channel")

	if err != nil {
		log.Fatalf("Unable to consume messages from %s due to: %s", queue, err.Error())
	}

	return msgs
}
