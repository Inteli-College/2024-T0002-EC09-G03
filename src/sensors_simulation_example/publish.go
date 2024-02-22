package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

// MyDataStructure represents the structure of the data to send.
type MyDataStructure struct {
	Name    string `json:"name"`
	Unit    string `json:"unit"`
	CoordsX string `json:"coords_x"`
	CoordsY string `json:"coords_y"`
	Date    string `json:"date"`
}

// DataWrapper wraps MyDataStructure for sending as a single "data" field.
type DataWrapper struct {
	Data MyDataStructure `json:"data"`
}

// failOnError logs any errors and exits.
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

	// Retrieve the RabbitMQ URL from environment variables
	rabbitMQURL := os.Getenv("RABBITMQ_URL") // Example: "amqp://user:password@localhost:5672/"

	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"MQTTSensors", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Prepare the data to send
	myData := MyDataStructure{
		Name:    "ExampleName",
		Unit:    "C",
		CoordsX: "23",
		CoordsY: "45",
		Date:    time.Now().Format(time.RFC3339),
	}

	// Wrap the data in DataWrapper
	dataWrapper := DataWrapper{
		Data: myData,
	}

	// Marshal the wrapped data into JSON
	body, err := json.Marshal(dataWrapper)
	failOnError(err, "Failed to marshal JSON")

	// Publish the message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	// Log the sent message
	log.Printf(" [x] Sent %s", body)
}
