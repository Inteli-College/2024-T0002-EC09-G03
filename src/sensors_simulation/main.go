package main

import (
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/generator"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
)

func init() {
	initialization.LoadEnvVariables()
}

func main() {
	sensorsInstanciated := generator.GenerateSensors(3, &sensors.Sensors)

	for _, sensor := range *sensorsInstanciated {
		go sensor.Emulate()
	}

	select {}
}
