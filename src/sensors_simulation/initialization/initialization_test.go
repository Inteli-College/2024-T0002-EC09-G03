package initialization

import (
	"os"
	"testing"
)

var envVariablesToCheck = [4]string{
	"BROKER_URL",
	"BROKER_PORT",
	"RABBIT_USER",
	"RABBIT_PASSWORD",
}

func TestEnvLoading(t *testing.T) {
	LoadEnvVariables("../.env")
	for _, envName := range envVariablesToCheck {
		if checking := os.Getenv(envName); checking == "" {
			t.Errorf("\033[31mExpected to load %s, but it wasn't able to read value.\033[0m", envName)
		}
	}
}
