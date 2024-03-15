package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorsData struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	SensorId  string             `bson:"sensor_id,omitempty" json:"id,omitempty"`
	Data      []DataSensor       `bson:"data" json:"data"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
