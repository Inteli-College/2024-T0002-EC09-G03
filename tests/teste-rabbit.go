package main

import (
	"os"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestRabbitMQ(t *testing.T) {
	// Obtém a URL do RabbitMQ da variável de ambiente
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		t.Fatal("A variável de ambiente RABBITMQ_URL não foi definida")
	}

	// Estabelece a conexão com o RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		t.Fatalf("Erro ao conectar ao RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Cria um canal na conexão
	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Erro ao criar o canal: %s", err)
	}
	defer ch.Close()

	// Declara a fila "sensor_data"
	_, err = ch.QueueDeclare(
		"sensor_data", // nome da fila
		true,          // durável
		false,         // exclusiva
		false,         // autoexcluir
		false,         // argumentos
	)
	if err != nil {
		t.Fatalf("Erro ao declarar a fila 'sensor_data': %s", err)
	}

	// Cria um consumidor para a fila "sensor_data"
	msgs, err := ch.Consume(
		"sensor_data", // nome da fila
		"",            // nome do consumidor
		true,          // automático
		false,         // exclusivo
		false,         // no-local
		false,         // no-wait
		nil,           // argumentos
	)
	if err != nil {
		t.Fatalf("Erro ao criar o consumidor: %s", err)
	}

	// Espera por mensagens da fila "sensor_data" por 1 segundo
	time.Sleep(time.Second)

	// Verifica se alguma mensagem foi recebida
	if len(msgs) > 0 {
		t.Log("Conexão RabbitMQ estabelecida com sucesso")
	} else {
		t.Error("Falha ao estabelecer conexão RabbitMQ")
	}

	// Fecha o canal e a conexão
	ch.Close()
	conn.Close()
}
