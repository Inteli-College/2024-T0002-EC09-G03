package infra_test

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	"github.com/Inteli-College/2024-T0002-EC09-G03/tests/mocks"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewDBConnection(t *testing.T) {
	mockConnector := mocks.MockMongoDBConnector{
		ConnectFunc: func(uri string) (*mongo.Database, error) {
			// Simula uma conexão bem-sucedida retornando um &mongo.Database mockado
			return &mongo.Database{}, nil
		},
	}

	_, err := infra.NewDBConnectionUsing(&mockConnector) // Adaptar a função original para usar a interface
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
