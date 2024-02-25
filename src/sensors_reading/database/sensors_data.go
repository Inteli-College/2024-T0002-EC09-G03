package database

import (
	"fmt"
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

func CreateSensorsData(db *gorm.DB, sensor_id string, data string, coordinate_x float64, coordinate_y float64, createdAt time.Time) SensorsData {
	id := uuid.New()
	sensorData := SensorsData{Id: id.String(), Sensor_id: sensor_id, Data: data, Coordinate_x: coordinate_x, Coordinate_y: coordinate_y, CreatedAt: createdAt}
	result := db.Create(&sensorData)
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
	}
	return sensorData
}

// func CreateSensorsData(db *gorm.DB, sensor_name string, data string, coordinate_x float64, coordinate_y float64, createdAt time.Time) SensorsData {
// 	var sensor Sensor = GetSensorOrCreate(db, sensor_name)
// 	id := uuid.New()
// 	sensorData := SensorsData{Id: id.String(), Sensor_id: sensor.Id, Data: data, Coordinate_x: coordinate_x, Coordinate_y: coordinate_y, CreatedAt: createdAt}
// 	result := db.Create(&sensorData)
// 	if result.Error != nil {
// 		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
// 	}
// 	return sensorData
// }

type SensorsReading struct {
	Id      string
	Name    string
	Data    string
	CoordsX float64
	CoordsY float64
	Date    time.Time
}

// func createSensorDataBatch(db *gorm.DB, sensors_reading_batch []SensorsReading) {
// 	var sensor Sensor = GetSensorOrCreate(db)
// 	id := uuid.New()
// 	sensorData := []SensorsData{Id: id.String(), Sensor_id: sensor.Id, Data: data, Coordinate_x: coordinate_x, Coordinate_y: coordinate_y, CreatedAt: createdAt}
// 	result := db.Create(&sensors_reading_batch)
// 	if result.Error != nil {
// 		panic("failed to create sensor data")
// 	}
// }
