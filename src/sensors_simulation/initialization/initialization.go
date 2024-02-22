package initialization

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var variablesToCheck = [4]string{
	"BROKER_URL",
	"BROKER_PORT",
	"RABBIT_USER",
	"RABBIT_PASSWORD",
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
