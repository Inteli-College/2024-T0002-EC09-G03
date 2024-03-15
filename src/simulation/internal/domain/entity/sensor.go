package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type SensorCallbackFunc func() (*[]DataSensor, error)

type Sensor struct {
	Name              string             `bson:"name"`
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	CoordinateX       float64            `bson:"coordinate_x"`
	CoordinateY       float64            `bson:"coordinate_y"`
	emulationCallback SensorCallbackFunc
}

func (s *Sensor) GetData() (*[]DataSensor, error) {
	return s.emulationCallback()
}

func (s *Sensor) SetCallback(callback SensorCallbackFunc) {
	s.emulationCallback = callback
}

func (s *Sensor) RemoveId() interface{} {
	return struct {
		Name        string  `bson:"name"`
		CoordinateX float64 `bson:"coordinate_x"`
		CoordinateY float64 `bson:"coordinate_y"`
	}{
		Name:        s.Name,
		CoordinateX: s.CoordinateX,
		CoordinateY: s.CoordinateY,
	}
}
