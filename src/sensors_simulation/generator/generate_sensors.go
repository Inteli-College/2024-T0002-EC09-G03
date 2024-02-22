package generator

import (
	"math/rand"
	"reflect"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors"
)

var existingCoords = make(map[string]map[float64]bool)

func generateCoords(sensor sensors.SensorInterface) (float64, float64) {
	baseX := -46.633308
	baseY := 23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		if existingCoords[sensor.GetSensorName()] == nil {
			existingCoords[sensor.GetSensorName()] = make(map[float64]bool)
		}

		if existingCoords[sensor.GetSensorName()][coordsX+coordsY] {
			continue
		} else {
			existingCoords[sensor.GetSensorName()][coordsX+coordsY] = true
			return coordsX, coordsY
		}
	}
}

func GenerateSensors(number int, sensorsTypes *[]reflect.Type) *[]interface{} {
	var instances []interface{}

	for _, t := range *sensorsTypes {
		for i := 0; i < number/len(*sensorsTypes); i++ {
			instance := reflect.New(t).Interface()

			if accertedInstance, ok := instance.(sensors.SensorInterface); ok {
				accertedInstance.New()
				accertedInstance.SetCoords(generateCoords(accertedInstance))
			}

			instances = append(instances, instance)
		}
	}

	return &instances
}
