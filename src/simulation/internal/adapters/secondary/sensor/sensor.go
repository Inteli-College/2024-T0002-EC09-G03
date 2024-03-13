package sensor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SensorAdapter struct {
	db *mongo.Database
}

func NewSensorDataAdapter(db *mongo.Database) *SensorAdapter {
	return &SensorAdapter{
		db: db,
	}
}

func (s *SensorAdapter) GetAllSensors() *[]entity.Sensor {

	coll := s.db.Collection("sensors")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	var sensors []entity.Sensor
	if err = cursor.All(context.TODO(), &sensors); err != nil {
		panic(err)
	}

	// Prints the sensors of the find operation as structs
	for _, result := range sensors {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	return &sensors
}

func (s *SensorAdapter) CreateSensor(sensor *entity.Sensor) (*entity.Sensor, error) {
	_, err := s.db.Collection("sensors").InsertOne(context.TODO(), sensor)
	if err != nil {
		panic("failed to create sensor")
	}
	fmt.Println("Sensor created")

	return sensor, nil
}
