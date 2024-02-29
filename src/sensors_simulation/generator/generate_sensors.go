package generator

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
	"gorm.io/gorm"
)

var existingCoords = make(map[string]map[float64]bool)
var mMap = sync.RWMutex{}
var mArr = sync.Mutex{}

func generateCoords(typeName *string) (*float64, *float64) {
	baseX := -46.633308
	baseY := -23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	defer mMap.Unlock()

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		mMap.Lock()
		if existingCoords[*typeName] == nil {
			existingCoords[*typeName] = make(map[float64]bool)
		}
		mMap.Unlock()

		mMap.RLock()
		if existingCoords[*typeName][coordsX+coordsY] {
			mMap.RUnlock()
			continue
		} else {
			mMap.RUnlock()
			mMap.Lock()
			existingCoords[*typeName][coordsX+coordsY] = true
			return &coordsX, &coordsY
		}
	}
}

func GenerateSensors(number int, sensorsTypes *map[string]sensors.SensorsTypes, db *gorm.DB, wg *sync.WaitGroup) *[]*sensors.Sensor {
	var instances []*sensors.Sensor

	for key, value := range *sensorsTypes {
		wg.Add(1)
		go func(key *string, value *sensors.SensorsTypes, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < number/len(*sensorsTypes); i++ {
				wg.Add(1)
				go func(value *sensors.SensorsTypes, wg *sync.WaitGroup) {
					defer wg.Done()
					coordsX, coordsY := generateCoords(key)
					log.Printf("Generating sensor: %s", fmt.Sprintf("%s %f,%f", value.Name, *coordsY, *coordsX))
					tempSensor := sensors.Sensor{}
					sensorName := fmt.Sprintf("%s %f,%f", value.Name, *coordsY, *coordsX)
					tempSensor.New(&sensorName, coordsX, coordsY, value.Callback, db)
					mArr.Lock()
					instances = append(instances, &tempSensor)
					mArr.Unlock()
				}(value, wg)
			}
		}(&key, &value, wg)
	}

	wg.Wait()

	return &instances
}
