package repository

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/mqtt"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: change this for the entity SendData
type SendData struct {
	Id        primitive.ObjectID  `json:"id"`
	Data      []entity.DataSensor `json:"data"`
	CreatedAt time.Time           `json:"created_at"`
}

type SensorRepo struct {
	sensorAdapter ports.SensorPort
}

func (s *SensorRepo) CreateSensor(name string, coordsX float64, coordsY float64, areaNumber int, callback entity.SensorCallbackFunc) (*entity.Sensor, ports.MQTTPort, error) {
	temp := &entity.Sensor{
		Name:        name,
		CoordinateX: coordsX,
		CoordinateY: coordsY,
		AreaNumber:  areaNumber,
	}

	temp.SetCallback(callback)

	sensor, err := s.sensorAdapter.CreateSensor(temp)

	mqttClient := &mqtt.MQTTAdapter{}

	mqttClient.CreateClient(&name)

	return sensor, mqttClient, err
}

func (s *SensorRepo) ActivateExistingSensor(id primitive.ObjectID, name string, coordsX float64, coordsY float64, callback entity.SensorCallbackFunc) (*entity.Sensor, ports.MQTTPort, error) {
	sensor := &entity.Sensor{
		Name:        name,
		Id:          id,
		CoordinateX: coordsX,
		CoordinateY: coordsY,
	}

	sensor.SetCallback(callback)

	mqttClient := &mqtt.MQTTAdapter{}

	mqttClient.CreateClient(&name)

	return sensor, mqttClient, nil
}

func (s *SensorRepo) randSleep() {
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
}

func (s *SensorRepo) EmulateSensor(sensor *entity.Sensor, mqttClient interface{}) {

	if mqttClient == nil {
		panic("Mqttclient undefined\n")
	}

	mqttAdapter, ok := mqttClient.(ports.MQTTPort)

	log.Printf("MQTT Adapter: %#v\n", mqttAdapter)

	if !ok {
		log.Printf("Sensor %#v failed to emulate: mqttClient undefined\n", *sensor)
	}
	for {
		// s.randSleep()
		simulatedData, err := sensor.GetData()

		if err != nil {
			log.Printf("Simulator %s couldn't simulate due to: %s\n", sensor.Name, err.Error())
			continue
		}

		sendData := SendData{
			Id:        sensor.Id,
			Data:      *simulatedData,
			CreatedAt: time.Now(),
		}

		// log.Printf("Sensor %s simulated data: %#v\n", sensor.Name, sendData)

		payload, _ := json.Marshal(sendData)

		topic := "sensor/data"
		mqttAdapter.Publish(&topic, 1, false, payload)
	}
}

func NewSensorRepo(sensorAdapter ports.SensorPort) *SensorRepo {
	return &SensorRepo{
		sensorAdapter: sensorAdapter,
	}
}
