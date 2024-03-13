package repository

import (
	"encoding/json"
	"log"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/google/uuid"
)

type SensorDataOnReceive struct {
	entity.SensorsData
	Data []entity.DataSensor `json:"data"`
}

func NewSensorData(info *[]byte) (*entity.SensorsData, error) {
	var temp SensorDataOnReceive

	err := json.Unmarshal(*info, &temp)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err)
		return nil, err
	}

	temp.SensorId = temp.Id
	temp.Id = uuid.New().String()

	sensorData, err := formatData(&temp)
	return sensorData, err
}

func formatData(sensorData *SensorDataOnReceive) (*entity.SensorsData, error) {
	dataToJson := struct {
		Data []entity.DataSensor `json:"data"`
	}{
		Data: sensorData.Data,
	}
	jsonData, err := json.Marshal(dataToJson)

	return &entity.SensorsData{
		Id:        sensorData.Id,
		SensorId:  sensorData.SensorId,
		Data:      string(jsonData),
		CreatedAt: sensorData.CreatedAt,
	}, err
}
