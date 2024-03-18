package infra

import (
	"time"

	mqttPaho "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockMongoDBConnector é um mock da interface MongoDBConnector.
type MockMongoDBConnector struct {
	ConnectFunc func(uri string) (*mongo.Database, error)
}

// Connect implementa a interface MongoDBConnector.
func (m MockMongoDBConnector) Connect(uri string) (*mongo.Database, error) {
	return m.ConnectFunc(uri)
}

// MockMQTTClient é um mock da interface mqttPaho.Client.
type MockMQTTClient struct {
	ConnectFunc func() mqttPaho.Token
	WaitFunc    func() error
}

// Connect implementa a interface mqttPaho.Client.
func (m MockMQTTClient) Connect() mqttPaho.Token {
	return m.ConnectFunc()
}

// Wait implementa a interface mqttPaho.Token.
func (m MockMQTTClient) Wait() bool {
	return m.WaitFunc() == nil
}

// WaitTimeout implementa a interface mqttPaho.Token.
func (m MockMQTTClient) WaitTimeout(d time.Duration) bool {
	return m.WaitFunc() == nil
}

// Error implementa a interface mqttPaho.Token.
func (m MockMQTTClient) Error() error {
	return m.WaitFunc()
}

// MockRabbitMQConnection é um mock da interface amqp.Connection.
type MockRabbitMQConnection struct {
	DialFunc func(url string) (*amqp.Connection, error)
}

// Dial implementa a interface amqp.Connection.
func (m MockRabbitMQConnection) Dial(url string) (*amqp.Connection, error) {
	return m.DialFunc(url)
}

// MockRabbitMQChannel é um mock da interface amqp.Channel.
type MockRabbitMQChannel struct {
	ConsumeFunc func(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

// Consume implementa a interface amqp.Channel.
func (m MockRabbitMQChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return make(<-chan amqp.Delivery), nil
}
