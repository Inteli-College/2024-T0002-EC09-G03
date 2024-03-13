package infra

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewQueueAdapter() *QueueAdapter {
	queue := QueueAdapter{
		msgsCh: make(map[string]<-chan amqp.Delivery),
	}

	queue.connect(os.Getenv("RABBITMQ_URL"))

	return &queue
}

type QueueAdapter struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	msgsCh map[string]<-chan amqp.Delivery
}

func (s *QueueAdapter) connect(url string) {
	conn, err := amqp.Dial(url)

	if err != nil {
		log.Fatalf("Fail to connect to rabbitmq due to: %s\n", err.Error())
	}

	log.Println("\033[32mRabbitMQ connected successfully\033[0m")

	s.conn = conn

	s.openChannel()
}

func (s *QueueAdapter) openChannel() {
	ch, err := s.conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open the channel due to: %s\n", err.Error())
	}

	log.Println("\033[32mRabbitMQ channel opened successfully\033[0m")

	s.ch = ch
}

func (s *QueueAdapter) GenerateConsumer(queue string) {
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

	s.msgsCh[queue] = msgs
}

func (s *QueueAdapter) RetriveLastMessage(queue string) *[]byte {
	msgs, ok := s.msgsCh[queue]

	if !ok {
		log.Fatalf("Unable to retrive channel of msgs for queue: %s\n", queue)
	}

	if msgs == nil {
		log.Fatalf("Channel doesn't exists for queue: %s\n", queue)
	}

	msg, _ := <-msgs

	return &msg.Body
}
