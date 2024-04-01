package infra

import (
	"context"
	"fmt"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestQueueLoad(t *testing.T) {
	connStr := "amqp://guest:guest@localhost:5672/"
	queueName := "your_queue_name_here"
	messageCount := 1000 // Ajuste conforme necessário

	conn, err := amqp.Dial(connStr)
	if err != nil {
		t.Fatalf("Could not establish connection with RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Could not open a channel: %v", err)
	}
	defer ch.Close()

	// Criando um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	startTime := time.Now()

	for i := 0; i < messageCount; i++ {
		body := fmt.Sprintf("Message %d", i)
		err = ch.PublishWithContext(
			ctx,
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			t.Errorf("Failed to publish message %d: %v", i, err)
			return
		}
	}

	elapsedTime := time.Since(startTime)
	t.Logf("%d messages sent in %s", messageCount, elapsedTime)
}
