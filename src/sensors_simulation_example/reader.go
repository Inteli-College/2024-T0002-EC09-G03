package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

// MyDataStructure matches the structure being sent by the publisher.
type MyDataStructure struct {
	Name    string `json:"name"`
	Unit    string `json:"unit"`
	CoordsX string `json:"coords_x"`
	CoordsY string `json:"coords_y"`
	Date    string `json:"date"`
}

// DataWrapper is used to unwrap the "data" field from the message.
type DataWrapper struct {
	Data MyDataStructure `json:"data"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rabbitMQURL := os.Getenv("RABBITMQ_URL") // Example: "amqp://consumerSensor:ConsumerSensorPassword@localhost:5672/"

	conn, err := amqp.Dial(rabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MQTTSensors", // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var wrapper DataWrapper
			err := json.Unmarshal(d.Body, &wrapper)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			log.Printf("Received a message: %+v", wrapper.Data)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
