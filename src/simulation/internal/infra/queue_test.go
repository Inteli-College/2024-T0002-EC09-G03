package infra

import (
	"testing"
)

// setUp cria os recursos necessários antes de cada teste
func setUp(t *testing.T) *QueueAdapter {
	t.Helper()

	queueAdapter := NewQueueAdapter()
	_, err := queueAdapter.ch.QueueDeclare(
		"testQueue", // name
		false,       // durable
		true,        // delete when unused
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
}

func TestGenerateConsumer(t *testing.T) {
	queueAdapter := setUp(t)
	defer tearDown(t, queueAdapter)

	queueAdapter.GenerateConsumer("testQueue")

	if _, ok := queueAdapter.msgsCh["testQueue"]; !ok {
		t.Fatalf("Failed to generate consumer for testQueue")
	}
}
