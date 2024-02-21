package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type myDataStructure struct {
	Data string `json:"data"`
}

func connectRabbitMQ() <-chan amqp.Delivery {
	conn, err := amqp.Dial("amqp://consumerSensor:ConsumerSensorPassword@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		defer conn.Close()
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		defer ch.Close()
	}

	messages, _ := ch.Consume(
		"MQTTSensors", // queue
		"",            // consumer
		true,          // auto-acknowledge
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	return messages
}

func main() {
	msgs := connectRabbitMQ()

	go func() {
		for income := range msgs {
			var data myDataStructure
			err := json.Unmarshal(income.Body, &data)

			if err != nil {
				log.Fatal("Not able to parser json")
			}
			log.Println(data.Data)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	select {}
}
