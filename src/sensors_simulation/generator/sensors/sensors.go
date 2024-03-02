package sensors

import (
	"fmt"
	"log"
	"strings"
	"sync"

	sensorTable "github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/database/sensor"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/generator/coords"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
)

var mArr = sync.Mutex{}

type dbInterface interface {
	CreateSensor(name *string, coordsX *float64, coordsY *float64) sensorTable.Sensor
	GetAllSensors() *[]*sensorTable.Sensor
}

func GenerateSensors(number int, sensorsTypes *map[string]sensors.SensorsTypes, db dbInterface, wg *sync.WaitGroup) *[]*sensors.Sensor {
	var instances []*sensors.Sensor

	existingSensors := db.GetAllSensors()
	amount := len(*existingSensors)
	log.Printf("Amount: %d\n", amount)
	log.Printf("Number: %d\n", number)

	switch {
	case amount >= number:
		log.Println("Sensors already exists, activating them...")
		activateSensors(existingSensors, sensorsTypes, db, wg, &instances)

	case amount == 0:
		log.Println("Creating all sensors from scratch...")
		createSensorsFromScratch(number, sensorsTypes, db, wg, &instances, nil)

	case amount < number && amount != 0:
		log.Println("Not enought sensors, creating more...")
		missing := number - amount
		log.Printf("Missing: %d\n", missing)
		createAndActivateSensors(missing, existingSensors, sensorsTypes, db, wg, &instances)

	default:
		panic("Sensors creation failed!")
	}

	wg.Wait()

	return &instances
}

func createSensorsFromScratch(number int, sensorsTypes *map[string]sensors.SensorsTypes, db dbInterface, wg *sync.WaitGroup, instances *[]*sensors.Sensor, sensorsExists *map[string]bool) {

	for key, value := range *sensorsTypes {
		wg.Add(1)
		maxIt := number / len(*sensorsTypes)

		createSensor := createSensorWithExistingCheck(sensorsExists, value, key, db, instances)

		go manageCreationPerType(wg, maxIt, createSensor)

	}
}

func manageCreationPerType(wg *sync.WaitGroup, maxIt int, createSensor func(int, *sync.WaitGroup)) {

	defer wg.Done()

	for i := 0; i < maxIt; i++ {
		wg.Add(1)

		go createSensor(i, wg)
	}
}

func createSensorWithExistingCheck(sensorsExists *map[string]bool, value sensors.SensorsTypes, key string, db dbInterface, instances *[]*sensors.Sensor) func(int, *sync.WaitGroup) {
	blockCheck := sync.Mutex{}

	return func(index int, wg *sync.WaitGroup) {
		defer wg.Done()
		sensorName := fmt.Sprintf("%s %d", value.Name, index)

		if sensorsExists != nil {
			for {
				blockCheck.Lock()
				if exists, ok := (*sensorsExists)[sensorName]; ok && exists {
					blockCheck.Unlock()
					index = index + 1
					sensorName = fmt.Sprintf("%s %d", value.Name, index)
					continue
				}
				(*sensorsExists)[sensorName] = true
				blockCheck.Unlock()
				break
			}
		}

		coordsX, coordsY := coords.GenerateCoords(&key)

		log.Printf("Generating sensor: %s\n", sensorName)
		tempSensor := sensors.New(&sensorName, coordsX, coordsY, value.Callback, db, wg)

		mArr.Lock()
		*instances = append(*instances, tempSensor)
		mArr.Unlock()
	}

}

func activateSensors(existingSensors *[]*sensorTable.Sensor, sensorsTypes *map[string]sensors.SensorsTypes, db dbInterface, wg *sync.WaitGroup, instances *[]*sensors.Sensor) *map[string]bool {

	sensorsCreated := map[string]bool{}

	for _, sensor := range *existingSensors {
		wg.Add(1)
		sensorsCreated[sensor.Name] = true

		sensorType := strings.Split(sensor.Name, " ")

		coords.MutexMap.Lock()
		if coords.ExistingCoords[sensorType[0]] == nil {
			coords.ExistingCoords[sensorType[0]] = make(map[float64]bool)
		}

		coords.ExistingCoords[sensorType[0]][sensor.CoordinateX+sensor.CoordinateY] = true
		coords.MutexMap.Unlock()

		go manageOneActivation(sensor, wg, sensorsTypes, db, instances)
	}

	return &sensorsCreated
}

func manageOneActivation(sensor *sensorTable.Sensor, wg *sync.WaitGroup, sensorsTypes *map[string]sensors.SensorsTypes, db dbInterface, instances *[]*sensors.Sensor) {
	defer wg.Done()
	parts := strings.SplitN(sensor.Name, " ", 2)
	if len(parts) == 0 {
		log.Fatalf("Error parsing sensor name: %s\n", sensor.Name)
	}
	sensorInfo, exists := (*sensorsTypes)[parts[0]]
	if !exists {
		log.Fatalf("Type %s doesn't exists\n", parts[0])
	}

	tempSensor := sensors.New(&sensor.Name, &sensor.CoordinateX, &sensor.CoordinateY, sensorInfo.Callback, db, wg, &sensor.Id)

	mArr.Lock()
	*instances = append(*instances, tempSensor)
	mArr.Unlock()
}

func createAndActivateSensors(number int, existingSensors *[]*sensorTable.Sensor, sensorsTypes *map[string]sensors.SensorsTypes, db dbInterface, wg *sync.WaitGroup, instances *[]*sensors.Sensor) {

	sensorsExists := activateSensors(existingSensors, sensorsTypes, db, wg, instances)
	createSensorsFromScratch(number, sensorsTypes, db, wg, instances, sensorsExists)

}
