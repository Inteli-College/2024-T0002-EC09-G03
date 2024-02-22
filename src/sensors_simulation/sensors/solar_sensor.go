package sensors

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SolarSensor struct {
	Name    string    `json:"name"`
	Unit    string    `json:"unit"`
	Data    string    `json:"data"`
	CoordsX float64   `json:"coords_x"`
	CoordsY float64   `json:"coords_y"`
	Date    time.Time `json:"date"`
}

func (s *SolarSensor) New() {
	s.Name = "sensor solar"
}

func (s *SolarSensor) GetSensorName() string {
	return s.Name
}

func (s *SolarSensor) SetCoords(coordsX float64, coordsY float64) {
	log.Println("Setando as coords")
	s.CoordsX = coordsX
	s.CoordsY = coordsY
}

func (s *SolarSensor) Emulate(mqqtClient *mqtt.Client) {

}

func (s *SolarSensor) setDateTimeNow() {
	s.Date = time.Now()
}

func (s *SolarSensor) generateReading() {

}
