package sensors

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/connections"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type simulationFunction func() ([]SensorData, error)

type Sensor struct {
	Name       string       `json:"name"`
	Data       []SensorData `json:"data"`
	CoordsX    float64      `json:"coords_x"`
	CoordsY    float64      `json:"coords_y"`
	Date       time.Time    `json:"date"`
	client     mqtt.Client
	simulation simulationFunction
}

type SensorData struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}

func (s *Sensor) New(name string, coordsX float64, coordsY float64, callback simulationFunction) {
	s.client = connections.GenerateClient(name)
	s.Name = name
	s.CoordsX = coordsX
	s.CoordsY = coordsY
	s.simulation = callback
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
		s.randSleep()
		simulatedData, err := s.simulation()

		if err != nil {
			log.Printf("Simulator %s couldn't simulate due to: %s\n", s.Name, err.Error())
			continue
		}

		s.Data = simulatedData

		payload, _ := json.Marshal(*s)

		s.setDateTimeNow()
		if token := s.client.Publish("sensor/data", 1, false, payload); token.Wait() && token.Error() != nil {

			log.Fatalf("From: %s - error: %s", s.Name, token.Error().Error())
		}
	}
}

type SensorsTypes struct {
	Name     string
	Unit     string
	Callback simulationFunction
}

var Sensors = map[string]SensorsTypes{
	"solar": {
		Name:     "SOLAR",
		Callback: SolarSimulation,
	},
	"gas": {
		Name:     "GAS",
		Callback: GasSimulation,
	},
	"polution": {
		Name:     "POLUTION",
		Callback: PolutionSimulation,
	},
}
