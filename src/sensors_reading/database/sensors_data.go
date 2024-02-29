package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type SensorsData struct {
	Id        string
	Sensor_id string
	Data      string
	CreatedAt time.Time
}

func CreateSensorsData(db *gorm.DB, sensor_id string, data string, coordinate_x float64, coordinate_y float64, createdAt time.Time) SensorsData {
	id := uuid.New()
	sensorData := SensorsData{Id: id.String(), Sensor_id: sensor_id, Data: data, CreatedAt: createdAt}
	result := db.Create(&sensorData)
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
	}
	return sensorData
}

func CreateSensorsDataBatch(db *gorm.DB, data *[]*SensorsData) {
	result := db.CreateInBatches(*data, len(*data))
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create sensor data: %s\n", result.Error.Error()))
	}
}

type SensorsReading struct {
	Id      string
	Name    string
	Data    string
	CoordsX float64
	CoordsY float64
	Date    time.Time
}
