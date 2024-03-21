package mqtt

import (
	"sync"
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	adapter := &MQTTAdapter{}
	clientName := "testClient"

	// Simula a criação de um cliente
	adapter.CreateClient(&clientName)

	assert.NotNil(t, adapter.client, "O cliente deve ser não nulo após a criação")
	assert.Equal(t, &clientName, adapter.clientName, "Os nomes dos clientes devem ser iguais")
}

func TestPublish(t *testing.T) {
	adapter := &MQTTAdapter{}
	clientName := "testClient"
	topic := "test/topic"
	payload := "Test payload"

	wg := sync.WaitGroup{}
	// Aqui, substitua infra.NewMQTTConnection pelo método real que cria a conexão com o broker MQTT
	adapter.client = infra.NewMQTTConnetion(&clientName, &wg)

	// Tentativa de publicar no tópico
	assert.NotPanics(t, func() {
		adapter.Publish(&topic, 0, false, payload)
	}, "A publicação não deve causar pânico")

	// Adicional: Verificar se a mensagem foi recebida por um subscriber, se necessário
}
