package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func New() *rabbitMQ {
	rabbit := rabbitMQ{
		url: os.Getenv("RABBITMQ_URL"),
	}

	rabbit.connect()
	rabbit.openChannel()

	return &rabbit
}

type rabbitMQ struct {
	url        string
	connection *amqp.Connection
	ch         *amqp.Channel
	msgs       <-chan amqp.Delivery
}

func (s *rabbitMQ) connect() {
	conn, err := amqp.Dial(s.url)

	if err != nil {
		log.Fatalf("Fail to connect to rabbitmq due to: %s\n", err.Error())
	}

	log.Println("\033[32mRabbitMQ connected successfully\033[0m")

	s.connection = conn
}

func (s *rabbitMQ) openChannel() {
	ch, err := s.connection.Channel()

	if err != nil {
		log.Fatalf("Failed to open the channel due to: %s\n", err.Error())
	}

	log.Println("\033[32mRabbitMQ channel opened successfully\033[0m")

	s.ch = ch
}

func (s *rabbitMQ) GenerateConsumer(queue string) {
	msgs, err := s.ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		log.Fatalf("Unable to consume messages from %s due to: %s", queue, err.Error())
	}

	log.Println("\033[32mRabbitMQ consumer generated successfully\033[0m")

	s.msgs = msgs
}

func (s *rabbitMQ) RetriveLastMessage() *[]byte {
	msg, ok := <-s.msgs

	if !ok {
		log.Fatalln("Unable to retrive msg from queue.")
	}

	return &msg.Body
}
