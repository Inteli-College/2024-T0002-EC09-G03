package sensors

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"
	"time"

	mqttConn "github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/connections/mqtt"
	sensorTable "github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/database/sensor"
)

type simulationFunction func() ([]SensorData, error)

type dbInterface interface {
	CreateSensor(name *string, coordsX *float64, coordsY *float64) sensorTable.Sensor
}

type mqttInterface interface {
	Publish(topic *string, qos byte, retained bool, payload interface{})
}

type SensorsTypes struct {
	Name     string
	Unit     string
	Callback simulationFunction
}

var Sensors = map[string]SensorsTypes{
	"Solar": {
		Name:     "Solar",
		Callback: SolarSimulation,
	},
	"Gas": {
		Name:     "Gas",
		Callback: GasSimulation,
	},
	"Polution": {
		Name:     "Polution",
		Callback: PolutionSimulation,
	},
}

type Sensor struct {
	Name       string       `json:"name"`
	Id         string       `json:"id"`
	Data       []SensorData `json:"data"`
	CoordsX    float64      `json:"coords_x"`
	CoordsY    float64      `json:"coords_y"`
	Date       time.Time    `json:"date"`
	client     mqttInterface
	simulation simulationFunction
}

type SensorData struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}

func (s *Sensor) setDateTimeNow() {
	s.Date = time.Now()
}

func (s *Sensor) randSleep() {
	sleepTime := rand.Intn(5) + 1
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

func (s *Sensor) Emulate() {
	for {
		// s.randSleep()
		simulatedData, err := s.simulation()

		if err != nil {
			log.Printf("Simulator %s couldn't simulate due to: %s\n", s.Name, err.Error())
			continue
		}

		s.Data = simulatedData

		payload, _ := json.Marshal(*s)

		s.setDateTimeNow()
		topic := "sensor/data"
		s.client.Publish(&topic, 1, false, payload)
	}
}

func New(name *string, coordsX *float64, coordsY *float64, callback simulationFunction, db dbInterface, wg *sync.WaitGroup, opts ...*string) *Sensor {

	var sensorId string

	if len(opts) == 0 {
		sensorId = db.CreateSensor(name, coordsX, coordsY).Id
	} else {
		sensorId = *opts[0]
	}
	client := mqttConn.New(name, wg)

	return &Sensor{
		Id:         sensorId,
		Name:       *name,
		CoordsX:    *coordsX,
		CoordsY:    *coordsY,
		simulation: callback,

		client: client,
	}

}
