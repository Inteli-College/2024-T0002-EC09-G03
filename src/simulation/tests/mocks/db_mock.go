package mocks

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MockMongoDBConnector é um mock da interface MongoDBConnector
type MockMongoDBConnector struct {
	ConnectFunc func(uri string) (*mongo.Database, error)
}

// Connect simula a conexão ao banco de dados MongoDB
func (m MockMongoDBConnector) Connect(uri string) (*mongo.Database, error) {
	return m.ConnectFunc(uri)
}
