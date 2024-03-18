package infra

import (
	"sync"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	mqttPaho "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TestNewDBConnection(t *testing.T) {
	// Cria um mock da interface MongoDBConnector.
	mockConnector := MockMongoDBConnector{
		// Define a função Connect para retornar um banco de dados mockado.
		ConnectFunc: func(uri string) (*mongo.Database, error) {
			return &mongo.Database{}, nil
		},
	}

	// Chama a função NewDBConnection com o mock como argumento.
	db, err := NewDBConnection(mockConnector)

	// Verifica se a função retornou um banco de dados válido e nenhum erro.
	if err != nil {
		t.Fatalf("Erro ao criar a conexão com o banco de dados: %v", err)
	}

	if db == nil {
		t.Fatalf("O banco de dados não foi criado")
	}
}

func TestNewMQTTConnection(t *testing.T) {
	// Cria um mock da interface mqttPaho.Client.
	mockClient := MockMQTTClient{
		// Define a função Connect para retornar um token mockado.
		ConnectFunc: func() mqttPaho.Token {
			return MockMQTTToken{}
		},
		// Define a função Wait para retornar nil, indicando uma conexão bem-sucedida.
		WaitFunc: func() error {
			return nil
		},
	}

	// Cria um WaitGroup para controlar a concorrência.
	wg := &sync.WaitGroup{}

	// Chama a função NewMQTTConnection com o mock como argumento.
	client := NewMQTTConnection("clientName", wg, &mockClient)

	// Verifica se a função retornou um cliente MQTT válido.
	if client == nil {
		t.Fatalf("O cliente MQTT não foi criado")
	}
}

func TestNewQueueAdapter(t *testing.T) {
	// Cria um mock da interface amqp.Connection.
	mockConnection := MockRabbitMQConnection{
		// Define a função Dial para retornar uma conexão mockada.
		DialFunc: func(url string) (*amqp.Connection, error) {
			return &amqp.Connection{}, nil
		},
	}

	// Cria um mock da interface amqp.Channel.
	mockChannel := MockRabbitMQChannel{
		// Define a função Consume para retornar um canal mockado.
		ConsumeFunc: func(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
			return make(<-chan amqp.Delivery), nil
		},
	}

	// Chama a função NewQueueAdapter com os mocks como argumentos.
	adapter := NewQueueAdapter(&mockConnection, &mockChannel)

	// Verifica se a função retornou um adaptador de fila válido.
	if adapter == nil {
		t.Fatalf("O adaptador de fila não foi criado")
	}

	// Verifica se a conexão com o RabbitMQ foi estabelecida.
	if adapter.conn == nil {
		t.Fatalf("A conexão com o RabbitMQ não foi estabelecida")
	}

	// Verifica se o canal do RabbitMQ foi criado.
	if adapter.ch == nil {
		t.Fatalf("O canal do RabbitMQ não foi criado")
	}
}
