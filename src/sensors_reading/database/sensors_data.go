package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *DB) CreateSensorsData(sensor_id string, data string, coordinate_x float64, coordinate_y float64, createdAt time.Time) {
	id := uuid.New()
	sensorData := SensorsData{Id: id.String(), Sensor_id: sensor_id, Data: data, CreatedAt: createdAt}
	result := s.db.Create(&sensorData)
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
	}
}

func (s *DB) CreateSensorsDataBatch(data *[]*SensorsData) {
	result := s.db.CreateInBatches(*data, len(*data))
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
	}
}
