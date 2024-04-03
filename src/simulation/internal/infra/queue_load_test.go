package infra

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TestQueueLoad(t *testing.T) {
	fmt.Printf("TESTANDO O LOAD QUEUE")

	path := "./.env"
	initialization.LoadEnvVariables([]string{"RABBITMQ_URL"}, &path)

	connStr := os.Getenv("RABBITMQ_URL")
	queueName := "testQueue"
	messageCount := 10 // Ajuste conforme necessário

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

	fmt.Printf("FIM DO TESTANDO O LOAD QUEUE")
}
