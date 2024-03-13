package entity

import (
	"time"
)

type SensorsData struct {
	Id        string `json:"id"`
	SensorId  string `json:"sensor_id"`
	Data      string `json:"data"`
	CreatedAt time.Time
}
