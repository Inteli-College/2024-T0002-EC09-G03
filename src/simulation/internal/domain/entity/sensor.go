package entity

type SensorCallbackFunc func() (*[]DataSensor, error)

type Sensor struct {
	Name              string  `json:"name"`
	Id                string  `json:"id"`
	CoordinateX       float64 `json:"coordinate_x"`
	CoordinateY       float64 `json:"coordinate_y"`
	emulationCallback SensorCallbackFunc
}

func (s *Sensor) GetData() (*[]DataSensor, error) {
	return s.emulationCallback()
}

func (s *Sensor) SetCallback(callback SensorCallbackFunc) {
	s.emulationCallback = callback
}
