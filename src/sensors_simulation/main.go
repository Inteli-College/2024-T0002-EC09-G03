package main

import (
	"github.com/Inteli-College/2024-T0002-EC09-G03/generator"
	"github.com/Inteli-College/2024-T0002-EC09-G03/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors"
)

func init() {
	initialization.LoadEnvVariables()
}

func main() {
	sensorsInstanciated := generator.GenerateSensors(100, &sensors.Sensors)

	for _, sensor := range *sensorsInstanciated {
		go sensor.Emulate()
	}

	select {}
}
