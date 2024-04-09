package infra

import (
	"fmt"
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
)

// setUp cria os recursos necessários antes de cada teste
func setUp(t *testing.T) *QueueAdapter {
	fmt.Printf("TESTANDO O QUEUE")
	t.Helper()

	path := "./.env"
	initialization.LoadEnvVariables([]string{"RABBITMQ_URL"}, &path)

	queueAdapter := NewQueueAdapter()
	_, err := queueAdapter.ch.QueueDeclare(
		"testQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		t.Fatalf("Unable to declare a test queue: %s", err)
	}

	return queueAdapter
}

// tearDown limpa os recursos após cada teste
func tearDown(t *testing.T, queueAdapter *QueueAdapter) {
	t.Helper()

	if int, err := queueAdapter.ch.QueueDelete("testQueue", false, false, false); err != nil {
		t.Fatalf("Unable to delete the test queue: %s, message: %d", err, int)
	}

	queueAdapter.conn.Close()
	fmt.Printf("FIM DO TESTANDO O QUEUE")
}

func TestGenerateConsumer(t *testing.T) {
	queueAdapter := setUp(t)
	defer tearDown(t, queueAdapter)

	queueAdapter.GenerateConsumer("testQueue")

	if _, ok := queueAdapter.msgsCh["testQueue"]; !ok {
		t.Fatalf("Failed to generate consumer for testQueue")
	}
}
