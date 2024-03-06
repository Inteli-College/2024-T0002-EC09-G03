package entity

import (
	"time"
)

type Sensor struct {
	Name string        `json:"name"`
	Id   string        `json:"id"`
	Data []SensorsData `json:"data"`
	Date time.Time     `json:"date"`
}
