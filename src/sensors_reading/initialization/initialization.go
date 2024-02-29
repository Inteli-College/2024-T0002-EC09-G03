package initialization

import (
	"errors"
	"fmt"
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

func checkEnvVarExistence(vars []string) error {
	for _, envName := range vars {
		if checking := os.Getenv(envName); checking == "" {
			return errors.New(fmt.Sprintf("Env variable %s not set. \n", envName))
		}
	}
	return nil
}

func LoadEnvVariables(path ...*string) {
	log.Print("Loading env variables...")

	var dotEnvErr error

	switch {
	case len(path) > 0:
		dotEnvErr = godotenv.Load(*path[0])
	default:
		dotEnvErr = godotenv.Load()
	}

	if errEnvMissing := checkEnvVarExistence(variablesToCheck[:4]); errEnvMissing != nil {

		if dotEnvErr != nil {
			log.Fatalf("%s\n%s\n", dotEnvErr.Error(), errEnvMissing.Error())
		}

		log.Fatalf("%s\n", errEnvMissing.Error())
	}

	log.Println("Environment variables loaded!")
}
