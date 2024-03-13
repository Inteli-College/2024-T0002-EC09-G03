package sensorData

import (
	"context"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type SensorDataAdapter struct {
	db *mongo.Database
}

func NewSensorDataAdapter(db *mongo.Database) *SensorDataAdapter {
	return &SensorDataAdapter{
		db: db,
	}
}

func (s *SensorDataAdapter) CreateInBatch(batch *[]*entity.SensorsData) error {
	batchNew := make([]interface{}, len(*batch))
	for i, v := range *batch {
		batchNew[i] = v
	}

	_, err := s.db.Collection("sensorsData").InsertMany(context.TODO(), batchNew)
	if err != nil {
		return err
	}

	return nil
}
