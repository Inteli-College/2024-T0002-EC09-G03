package reader

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Sensor struct {
	Name    string       `json:"name"`
	Data    []SensorData `json:"data"`
	CoordsX float64      `json:"coords_x"`
	CoordsY float64      `json:"coords_y"`
	Date    time.Time    `json:"date"`
}

type SensorData struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func Reader() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rabbitMQURL := os.Getenv("RABBITMQ_URL")

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
		"MQTTSensors", // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages from the queue
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
			var target Sensor
			err := json.Unmarshal(d.Body, &target)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
