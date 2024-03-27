package infra

import (
	"os"
	"sync"
	"testing"
)

// Substitua por seus próprios valores ou configure estes no ambiente de teste
func setupTestEnvironment() {
	os.Setenv("BROKER_URL", "localhost")
	os.Setenv("BROKER_PORT", "1883")
	os.Setenv("RABBIT_USER", "publisherSensor")
	os.Setenv("RABBIT_PASSWORD", "PublisherSensorPasswords")
}

func TestNewMQTTConnection(t *testing.T) {
	// setupTestEnvironment() // Deixando o setup de variaveis de ambiente para o github actions

	clientName := "testClient"
	wg := &sync.WaitGroup{}

	client := NewMQTTConnetion(&clientName, wg)

	if client == nil {
		t.Fatalf("Expected non-nil MQTT client")
	}

	if !client.IsConnected() {
		t.Fatalf("Expected MQTT client to be connected")
	}
}
