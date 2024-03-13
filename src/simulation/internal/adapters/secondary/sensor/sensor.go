package sensor

import (
	"fmt"
	"log"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"gorm.io/gorm"
)

type SensorAdapter struct {
	db *gorm.DB
}

func NewSensorDataAdapter(db *gorm.DB) *SensorAdapter {
	return &SensorAdapter{
		db: db,
	}
}

func (s *SensorAdapter) GetAllSensors() *[]entity.Sensor {
	var sensors []entity.Sensor

	err := s.db.Select("id", "name").Find(&sensors)

	if err.Error != nil {
		log.Fatalf("Error retriving existing data: %s\n", err.Error.Error())
	}

	return &sensors
}

func (s *SensorAdapter) CreateSensor(sensor *entity.Sensor) (*entity.Sensor, error) {
	result := s.db.Create(&sensor)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("failed to create sensor")
	}
	fmt.Println("Sensor created")

	return sensor, nil
}
