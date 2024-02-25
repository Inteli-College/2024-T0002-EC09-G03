package initialization

import (
	"errors"
	"fmt"
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

func checkEnvVarExistence(vars []string) error {
	for _, envName := range vars {
		if checking := os.Getenv(envName); checking == "" {
			return errors.New(fmt.Sprintf("Env variable %s not set. \n", envName))
		}
	}
	return nil
}

func LoadEnvVariables(path ...string) {
	log.Print("Loading .env variables...")

  if envLoadingErr := checkEnvVarExistence(variablesToCheck[:4]); envLoadingErr != nil {

		var finalPath string
		if len(path) > 0 && path[0] != "" {
			finalPath = path[0]
		} else {
			finalPath = "./.env"
		}

		err := godotenv.Load(finalPath)
		if err != nil {
			log.Fatalf("Error when loading .env file in the path: %s\n", finalPath)
		}

    if envMissing := checkEnvVarExistence(variablesToCheck[:4]); envMissing != nil {

			log.Fatalf("Env variable %s not set. \n", envMissing)
		}

		return
	}

	log.Println("Environment variables already loaded!")
}
