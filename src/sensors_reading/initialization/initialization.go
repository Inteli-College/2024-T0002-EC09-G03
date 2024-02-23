package initialization

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var variablesToCheck = [6]string{
	"RABBITMQ_URL",
	"DATABASE_HOST",
	"DATABASE_USER",
	"DATABASE_PASSWORD",
	"DATABASE_NAME",
	"DATABASE_PORT",
}

func LoadEnvVariables() {
	log.Print("Loading .env variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro carregando o arquivo .env")
	}

	for _, envName := range variablesToCheck {
		if checking := os.Getenv(envName); checking == "" {
			log.Fatalf("Env variable %s not set. \n", envName)
		}
	}

}
