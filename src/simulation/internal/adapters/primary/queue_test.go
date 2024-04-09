package primary

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	testRabbitMQQueue = "testQueue"
	testMongoDBCol    = "sensorsData"
	testMessage       = "Test sensor data message from RabbitMQ"
)

func TestMessageHandlerAdapter(t *testing.T) {

	path := "./.env"
	initialization.LoadEnvVariables([]string{"RABBITMQ_URL", "MONGODB_URI"}, &path)

	// Conexão com o MongoDB
	db, mongoClient := infra.NewDBConnection()
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Limpa a coleção de testes para garantir o ambiente
	if err := db.Collection(testMongoDBCol).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	// Conexão com o RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		testRabbitMQQueue, // queue name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Publica uma mensagem no RabbitMQ
	err = channel.PublishWithContext(
		context.Background(), // add context.Background() as the first argument
		"",                   // exchange
		q.Name,               // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(testMessage),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	// Espera para garantir que a mensagem seja processada
	time.Sleep(5 * time.Second)

	// Verifica no MongoDB para confirmar o recebimento e armazenamento da mensagem
	// var result bson.M
	// err = db.Collection(testMongoDBCol).FindOne(context.Background(), bson.M{"message": testMessage}).Decode(&result)
	// assert.Nil(t, err, "Error should be nil when fetching the processed message from MongoDB")
	// assert.NotEmpty(t, result, "The result should not be empty after processing the message")
}
