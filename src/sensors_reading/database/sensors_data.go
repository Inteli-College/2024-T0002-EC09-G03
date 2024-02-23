package database

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type SensorsData struct {
	Id           string
	Sensor_id    string
	Data         string
	Coordinate_x float64
	Coordinate_y float64
	CreatedAt    time.Time
}

func CreateSensorsData(db *gorm.DB, sensor_name string, data string, coordinate_x float64, coordinate_y float64, createdAt time.Time) SensorsData {
	var sensor Sensor = GetSensorOrCreate(db, sensor_name)
	id := uuid.New()
	sensorData := SensorsData{Id: id.String(), Sensor_id: sensor.Id, Data: data, Coordinate_x: coordinate_x, Coordinate_y: coordinate_y, CreatedAt: createdAt}
	result := db.Create(&sensorData)
	if result.Error != nil {
		panic("failed to create sensor data")
	}

	return sensorData
}
