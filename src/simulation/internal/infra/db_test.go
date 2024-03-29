package infra

import (
	"context"
	"os"
	"testing"
)

func TestNewDBConnection(t *testing.T) {
	os.Setenv("MONGODB_URI", "mongodb://root:password@localhost:27017/urbanpulsesp?retryWrites=true&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1")

	db, client := NewDBConnection()

	if db == nil {
		t.Fatal("Expected non-nil database connection")
	}

	if client == nil {
		t.Fatal("Expected non-nil database client")
	}

	// Opcional: Adicione aqui quaisquer verificações adicionais que desejar

	err := client.Disconnect(context.TODO())
	if err != nil {
		t.Fatalf("Failed to close database connection: %v", err)
	}
}
