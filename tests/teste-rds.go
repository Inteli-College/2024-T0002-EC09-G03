package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestRDS(t *testing.T) {
	// Obtém os parâmetros de conexão com o PostgreSQL das variáveis de ambiente
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")

	// Monta a string de conexão com o PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	// Estabelece a conexão com o PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Erro ao conectar ao PostgreSQL: %s", err)
	}
	defer db.Close()

	// Cria uma consulta para selecionar todos os registros da tabela "sensor_data"
	query := "SELECT * FROM sensor_data"

	// Executa a consulta
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("Erro ao executar a consulta: %s", err)
	}
	defer rows.Close()

	// Verifica se algum registro foi recuperado
	if rows.Next() {
		t.Log("Conexão RDS estabelecida com sucesso")
	} else {
		t.Error("Falha ao estabelecer conexão RDS")
	}
}
