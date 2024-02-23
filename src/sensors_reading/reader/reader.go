package reader

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
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

func Reader(db *gorm.DB) {
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

		var sensors_reading_batch [1000]Sensor
		count := 0
		for d := range msgs {
			var sensor Sensor
			err := json.Unmarshal(d.Body, &sensor)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				return
			}
			// log.Printf("Received a message: %+v", sensor)
			dataJson, _ := json.Marshal(sensor.Data)

			if count%1000 == 0 {
				database.CreateSensorsData(db, sensor.Name, string(dataJson), sensor.CoordsX, sensor.CoordsY, sensor.Date)
				count = 0
			}

			sensors_reading_batch[count] = sensor
			count += 1
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
