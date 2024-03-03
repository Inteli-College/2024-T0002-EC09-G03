package rabbitmq_test

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/connections/rabbitmq"
)

func TestGenerateConsumer(t *testing.T) {
	// Este teste requer que o RabbitMQ esteja rodando. Ele verifica se a função é capaz de criar e retornar uma instância válida (não nula) de um consumidor de mensagens do RabbitMQ.
	consumer := rabbitmq.GenerateConsumer("MQTTSensors")
	if consumer == nil {
		t.Error("expected non-nil consumer")
	}
}
