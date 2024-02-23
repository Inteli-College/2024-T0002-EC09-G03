package main

import (
	"os"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/generator"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
)

func init() {
	envPath := ""

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	initialization.LoadEnvVariables(envPath)
}

func main() {
	sensorsInstanciated := generator.GenerateSensors(3, &sensors.Sensors)

	for _, sensor := range *sensorsInstanciated {
		go sensor.Emulate()
	}

	select {}
}
