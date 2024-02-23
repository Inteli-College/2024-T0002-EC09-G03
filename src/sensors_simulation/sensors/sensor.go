package sensors

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/conections"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type simulationFunction func() (float64, error)

type Sensor struct {
	Name       string    `json:"name"`
	Unit       string    `json:"unit"`
	Data       float64   `json:"data"`
	CoordsX    float64   `json:"coords_x"`
	CoordsY    float64   `json:"coords_y"`
	Date       time.Time `json:"date"`
	client     mqtt.Client
	simulation simulationFunction
}

func (s *Sensor) New(name string, unit string, coordsX float64, coordsY float64, callback simulationFunction) {
	s.client = conections.GenerateClient(name)
	s.Name = name
	s.Unit = unit
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

		payload, _ := json.Marshal(s.Data)

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
		Unit:     "W/m^2",
		Callback: SolarSimulation,
	},
}

