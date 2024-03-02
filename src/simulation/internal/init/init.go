package init

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TODO: missing differentiation cmd
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
