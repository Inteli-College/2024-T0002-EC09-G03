package init

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func checkEnvVarExistence(vars []string) error {
	for _, envName := range vars {
		if checking := os.Getenv(envName); checking == "" {
			return errors.New(fmt.Sprintf("Env variable %s not set. \n", envName))
		}
	}
	return nil
}

func LoadEnvVariables(vars []string, path ...*string) {
	log.Print("Loading env variables...")

	var dotEnvErr error

	switch {
	case len(path) > 0:
		dotEnvErr = godotenv.Load(*path[0])
	default:
		dotEnvErr = godotenv.Load()
	}

	if errEnvMissing := checkEnvVarExistence(vars); errEnvMissing != nil {

		if dotEnvErr != nil {
			log.Fatalf("%s\n%s\n", dotEnvErr.Error(), errEnvMissing.Error())
		}

		log.Fatalf("%s\n", errEnvMissing.Error())
	}

	log.Println("Environment variables loaded!")
}
