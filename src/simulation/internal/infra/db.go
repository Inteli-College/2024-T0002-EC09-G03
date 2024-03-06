package infra

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic("DATABASE_PORT must be an integer")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Limite de tempo para considerar uma query "lenta"
			LogLevel:                  logger.Silent, // Define o nível de log
			IgnoreRecordNotFoundError: true,          // Ignora erros de 'registro não encontrado'
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	opts, _ := db.DB()

	opts.SetMaxOpenConns(98)
	opts.SetMaxIdleConns(2)
	opts.SetConnMaxLifetime(5 * time.Minute)

	return db
}
