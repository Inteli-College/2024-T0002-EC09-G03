package sensorData

import (
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"gorm.io/gorm"
)

type SensorDataAdapter struct {
	db *gorm.DB
}

func NewSensorDataAdapter(db *gorm.DB) *SensorDataAdapter {
	return &SensorDataAdapter{
		db: db,
	}
}

func (s *SensorDataAdapter) CreateInBatch(batch *[]*entity.SensorsData) error {
	result := s.db.CreateInBatches(*batch, len(*batch))
	if result.Error != nil {
		return result.Error
	}

	return nil
}
