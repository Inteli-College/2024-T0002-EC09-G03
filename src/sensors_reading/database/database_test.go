package tests

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// Esse teste verifica a capacidade de estabelecer uma conexão com o banco de dados utilizando o GORM
func TestDatabaseConnection(t *testing.T) {
	dsn := "host=localhost user=user_name password=password dbname=database_name port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
