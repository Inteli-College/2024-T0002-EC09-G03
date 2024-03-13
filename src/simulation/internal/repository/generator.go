package repository

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/ports"
)

type GeneratorRepo struct {
	sensorAdapter ports.SensorPort
	sensorRepo    *SensorRepo
	gen           *entity.Generator
}

func (s *GeneratorRepo) GenerateSensors(number int) *entity.SensorsInstances {
	existingSensors := s.sensorAdapter.GetAllSensors()
	amount := len(*existingSensors)
	log.Printf("Amount: %d\n", amount)
	log.Printf("Number: %d\n", number)

	switch {
	case amount >= number:
		log.Println("Sensors already exists, activating them...")
		s.activateSensors(existingSensors)

	case amount == 0:
		log.Println("Creating all sensors from scratch...")
		s.createSensorsFromScratch(number, nil)

	case amount < number && amount != 0:
		log.Println("Not enought sensors, creating more...")
		missing := number - amount
		log.Printf("Missing: %d\n", missing)
		s.createAndActivateSensors(missing, existingSensors)

	default:
		panic("Sensors creation failed!")
	}

	s.gen.Wg.Wait()

	return s.gen.Instances
}

func (s *GeneratorRepo) activateSensors(existingSensors *[]entity.Sensor) *map[string]bool {

	sensorsCreated := map[string]bool{}

	for _, sensor := range *existingSensors {
		s.gen.Wg.Add(1)
		sensorsCreated[sensor.Name] = true

		sensorSplited := strings.Split(sensor.Name, " ")
		if sensorSplited == nil || len(sensorSplited) == 0 {
			log.Fatalf("Error parsing sensor name: %s\n", sensor.Name)
		}

		if _, ok := s.gen.SensorTypes[sensorSplited[0]]; !ok {
			log.Fatalf("Type %s doesn't exists\n", sensorSplited[0])
		}

		s.gen.MMap.Lock()
		if s.gen.ExistingCoords[sensorSplited[0]] == nil {
			s.gen.ExistingCoords[sensorSplited[0]] = make(map[float64]bool)
		}

		s.gen.ExistingCoords[sensorSplited[0]][sensor.CoordinateX+sensor.CoordinateY] = true
		s.gen.MMap.Unlock()

		go s.manageOneActivation(sensor, s.gen.SensorTypes[sensorSplited[0]].Callback)
	}

	s.gen.Wg.Wait()

	return &sensorsCreated
}

func (s *GeneratorRepo) manageOneActivation(sensor entity.Sensor, sensorCallback entity.SensorCallbackFunc) {
	defer s.gen.Wg.Done()

	tempSensor, mqttClientInterface, err := s.sensorRepo.ActivateExistingSensor(sensor.Id, sensor.Name, sensor.CoordinateX, sensor.CoordinateY, sensorCallback)

	mqttClient, ok := mqttClientInterface.(ports.MQTTPort)
	if !ok {
		panic("Client does not implement ports.MQTTPort")
	}

	if err != nil {
		panic("Error activating existing sensors!\n")
	}

	s.gen.MArr.Lock()
	*s.gen.Instances = append(*s.gen.Instances, entity.SensorClient{Sensor: *tempSensor, Client: mqttClient})
	s.gen.MArr.Unlock()
}

func (s *GeneratorRepo) createSensorsFromScratch(number int, sensorsExists *map[string]bool) {

	for key, value := range s.gen.SensorTypes {
		s.gen.Wg.Add(1)
		maxIt := number / len(s.gen.SensorTypes)

		createSensor := s.createSensorWithExistingCheck(sensorsExists, value, key)

		go s.manageCreationPerType(maxIt, createSensor)

	}
}

func (s *GeneratorRepo) createSensorWithExistingCheck(sensorsExists *map[string]bool, value entity.SensorsTypes, key string) func(int, *sync.WaitGroup) {
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

		coordsX, coordsY := s.generateCoords(&key)

		log.Printf("Generating sensor: %s\n", sensorName)
		tempSensor, mqttClientInterface, err := s.sensorRepo.CreateSensor(sensorName, *coordsX, *coordsY, s.gen.SensorTypes[value.Name].Callback)

		mqttClient, ok := mqttClientInterface.(ports.MQTTPort)
		if !ok {
			panic("Client does not implement ports.MQTTPort")
		}

		if err != nil {
			panic("Error creating a new sensor with existing check")
		}

		s.gen.MArr.Lock()
		*s.gen.Instances = append(*s.gen.Instances, entity.SensorClient{Sensor: *tempSensor, Client: mqttClient})
		s.gen.MArr.Unlock()
	}

}

func (s *GeneratorRepo) manageCreationPerType(maxIt int, createSensor func(int, *sync.WaitGroup)) {

	defer s.gen.Wg.Done()

	for i := 0; i < maxIt; i++ {
		s.gen.Wg.Add(1)

		go createSensor(i, s.gen.Wg)
	}
}

func (s *GeneratorRepo) createAndActivateSensors(number int, existingSensors *[]entity.Sensor) {

	sensorsExists := s.activateSensors(existingSensors)
	s.createSensorsFromScratch(number, sensorsExists)

}

func (s *GeneratorRepo) generateCoords(typeName *string) (*float64, *float64) {
	baseX := -46.633308
	baseY := -23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	defer s.gen.MMap.Unlock()

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		s.gen.MMap.Lock()
		if s.gen.ExistingCoords[*typeName] == nil {
			s.gen.ExistingCoords[*typeName] = make(map[float64]bool)
		}
		s.gen.MMap.Unlock()

		s.gen.MMap.RLock()
		if s.gen.ExistingCoords[*typeName][coordsX+coordsY] {
			s.gen.MMap.RUnlock()
			continue
		} else {
			s.gen.MMap.RUnlock()
			s.gen.MMap.Lock()
			s.gen.ExistingCoords[*typeName][coordsX+coordsY] = true
			return &coordsX, &coordsY
		}
	}
}

func NewGeneratorRepo(sensorAdapter ports.SensorPort, generator *entity.Generator, sensorRepo *SensorRepo) *GeneratorRepo {
	return &GeneratorRepo{
		sensorAdapter: sensorAdapter,
		sensorRepo:    sensorRepo,
		gen:           generator,
	}
}
