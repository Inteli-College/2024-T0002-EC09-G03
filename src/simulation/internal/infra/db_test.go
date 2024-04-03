package infra

import (
	"context"
	"fmt"
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
)

func TestNewDBConnection(t *testing.T) {
	fmt.Printf("TESTANDO O DB")

	path := "./.env"
	initialization.LoadEnvVariables([]string{"MONGODB_URI"}, &path)

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
	fmt.Printf("FIM DO TESTANDO O DB")
}
