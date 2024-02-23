package database

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sensor struct {
	Id   string
	Name string
}

func GetSensorByName(db *gorm.DB, name string) Sensor {
	var sensor Sensor
	db.Where("name = ?", name).First(&sensor)
	return sensor
}

func CreateSensor(db *gorm.DB, name string) Sensor {
	id := uuid.New()
	sensor := Sensor{Id: id.String(), Name: name}
	result := db.Create(&sensor)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("failed to create sensor")
	}
	fmt.Println("Sensor created")

	return sensor
}

func GetSensorOrCreate(db *gorm.DB, name string) Sensor {
	sensor := GetSensorByName(db, name)
	if sensor.Name == "" {
		sensor = CreateSensor(db, name)
	}
	return sensor
}
