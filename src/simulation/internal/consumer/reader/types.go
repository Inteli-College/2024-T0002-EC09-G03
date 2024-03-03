package reader

import (
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/consumer/database"
)

type Sensor struct {
	Name string       `json:"name"`
	Id   string       `json:"id"`
	Data []SensorData `json:"data"`
	Date time.Time    `json:"date"`
}

type SensorData struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}

type dbInterface interface {
	CreateSensorsDataBatch(data *[]*database.SensorsData)
}

type queueInterface interface {
	RetriveLastMessage() *[]byte
}
