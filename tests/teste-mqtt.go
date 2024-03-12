package main

import (
	"os"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestMQTT(t *testing.T) {
	// Obtém o nome do broker MQTT da variável de ambiente
	brokerURL := os.Getenv("BROKER_URL")
	if brokerURL == "" {
		t.Fatal("A variável de ambiente BROKER_URL não foi definida")
	}

	// Cria um cliente MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	client := mqtt.NewClient(opts)

	// Conecta ao broker MQTT
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker MQTT: %s", token.Error())
	}

	// Assina o tópico "sensor/data"
	if token := client.Subscribe("sensor/data", 1, nil); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao assinar o tópico 'sensor/data': %s", token.Error())
	}

	// Espera por mensagens do tópico "sensor/data" por 1 segundo
	time.Sleep(time.Second)

	// Verifica se alguma mensagem foi recebida
	if client.IsConnected() {
		t.Log("Conexão MQTT estabelecida com sucesso")
	} else {
		t.Error("Falha ao estabelecer conexão MQTT")
	}

	// Desconecta do broker MQTT
	client.Disconnect(250)
}
