package generator

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
)

var existingCoords = make(map[string]map[float64]bool)

func generateCoords(typeName string) (float64, float64) {
	baseX := -46.633308
	baseY := 23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		if existingCoords[typeName] == nil {
			existingCoords[typeName] = make(map[float64]bool)
		}

		if existingCoords[typeName][coordsX+coordsY] {
			continue
		} else {
			existingCoords[typeName][coordsX+coordsY] = true
			return coordsX, coordsY
		}
	}
}

func GenerateSensors(number int, sensorsTypes *map[string]sensors.SensorsTypes) *[]*sensors.Sensor {
	var instances []*sensors.Sensor

	for key, value := range *sensorsTypes {
		for i := 0; i < number/len(*sensorsTypes); i++ {
      coordsX, coordsY := generateCoords(key)
      log.Printf("Generating sensor: %s", fmt.Sprintf("%s %f,%f", value.Name, coordsX, coordsY))
			tempSensor := sensors.Sensor{}
			tempSensor.New(fmt.Sprintf("%s %f,%f", value.Name, coordsX, coordsY), coordsX, coordsY, value.Callback)
			instances = append(instances, &tempSensor)
		}
	}

	return &instances
}
