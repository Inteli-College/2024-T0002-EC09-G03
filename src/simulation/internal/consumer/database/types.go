package database

import "time"

type SensorsData struct {
	Id        string
	Sensor_id string
	Data      string
	CreatedAt time.Time
}
