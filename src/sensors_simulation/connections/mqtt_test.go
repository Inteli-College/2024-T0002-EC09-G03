package connections

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/initialization"
)

func TestGenerateClient(t *testing.T) {
	initialization.LoadEnvVariables("../.env")

	client := GenerateClient("testing_client")

	if token := client.Publish("testing/test", 1, false, "Hello test"); token.Wait() && token.Error() != nil {
		t.Errorf("\033[31mClient not health. Wasn't able to publish due to: %s\n\033[0m", token.Error().Error())
	}

}
