package infra

import (
	"fmt"
	"sync"
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
)

func TestNewMQTTConnection(t *testing.T) {

	fmt.Printf("TESTANDO O MQTT")
	path := "./.env"
	initialization.LoadEnvVariables([]string{"BROKER_URL", "BROKER_PORT", "RABBIT_USER", "RABBIT_PASSWORD"}, &path)

	clientName := "testClient"
	wg := &sync.WaitGroup{}

	client := NewMQTTConnetion(&clientName, wg)

	if client == nil {
		t.Fatalf("Expected non-nil MQTT client")
	}

	if !client.IsConnected() {
		t.Fatalf("Expected MQTT client to be connected")
	}
	fmt.Printf("FIM DO TESTANDO O MQTT")
}
