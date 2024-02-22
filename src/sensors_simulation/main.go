package main

import (
	"log"
	"reflect"

	"github.com/Inteli-College/2024-T0002-EC09-G03/conections"
	"github.com/Inteli-College/2024-T0002-EC09-G03/generator"
	"github.com/Inteli-College/2024-T0002-EC09-G03/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors"
)

func init() {
	initialization.LoadEnvVariables()
}

func main() {
	mqttClient := conections.GenerateClient("sensors")

	var sensorsInstanciated = generator.GenerateSensors(1, &[]reflect.Type{
		reflect.TypeOf(sensors.SolarSensor{}),
	})

	for _, sensor := range *sensorsInstanciated {
		if s, ok := sensor.(sensors.SensorInterface); ok {
			go s.Emulate(&mqttClient)
		} else {
			log.Fatalf("Not able to assert sensor: %#v", sensor)
		}
	}

	select {}
}
