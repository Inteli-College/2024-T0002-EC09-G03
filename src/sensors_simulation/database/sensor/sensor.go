package sensor

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sensor struct {
	Id          string
	Name        string
	CoordinateX float64
	CoordinateY float64
}

func New(db *gorm.DB) *SensorGorm {
	return &SensorGorm{
		db: db,
	}
}

type SensorGorm struct {
	db *gorm.DB
}

func (s *SensorGorm) GetAllSensors() *[]*Sensor {
	var sensors []*Sensor

	err := s.db.Select("id", "name").Find(&sensors)

	if err.Error != nil {
		log.Fatalf("Error retriving existing data: %s\n", err.Error.Error())
	}

	return &sensors
}

func (s *SensorGorm) CreateSensor(name *string, coordsX *float64, coordsY *float64) Sensor {
	id := uuid.New()
	sensor := Sensor{Id: id.String(), Name: *name, CoordinateX: *coordsX, CoordinateY: *coordsY}

	result := s.db.Create(&sensor)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("failed to create sensor")
	}
	fmt.Println("Sensor created")

	return sensor
}
