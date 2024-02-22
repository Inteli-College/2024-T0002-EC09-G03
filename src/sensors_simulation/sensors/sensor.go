package sensors

import (
	// "time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sensor struct {
	Name    string `json:"name"`
	Unit    string `json:"unit"`
	Data    string `json:"data"`
	CoordsX string `json:"coords_x"`
	CoordsY string `json:"coords_y"`
	Date    string `json:"date"`
}

type SensorInterface interface {
	// generateReading()
  New()
  GetSensorName() string
  SetCoords(coordsX float64, coordsY float64)
	Emulate(mqqtClient *mqtt.Client)
	// publishReading(mqqtClient *mqtt.Client) error
	// getDateTime() time.Time
}
