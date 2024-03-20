package repository

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockSensorAdapter struct{}

func (m *mockSensorAdapter) GetAllSensors() *[]entity.Sensor {
	sensors := make([]entity.Sensor, 0)
	return &sensors
}

func (m *mockSensorAdapter) CreateSensor(sensor *entity.Sensor) (*entity.Sensor, error) {
	sensor.Id = primitive.NewObjectID()
	return sensor, nil
}

func TestCreateAndActivateSensor(t *testing.T) {
	sensorAdapter := &mockSensorAdapter{}
	repo := NewSensorRepo(sensorAdapter)
	name := "TestSensor"
	coordsX := 10.0
	coordsY := 20.0
	// Certifique-se de que o callback corresponda à assinatura esperada.
	callback := func() (*[]entity.DataSensor, error) {
		// Sua lógica aqui, por exemplo, criar um slice de DataSensor.
		dataSensors := make([]entity.DataSensor, 0)
		// Adicione dados ao dataSensors conforme necessário.

		return &dataSensors, nil
	}

	// Teste CreateSensor
	sensor, mqttClient, err := repo.CreateSensor(name, coordsX, coordsY, callback)
	assert.NoError(t, err)
	assert.NotNil(t, sensor)
	assert.NotNil(t, mqttClient)

	// Geração do ID para testar ActivateExistingSensor
	id := primitive.NewObjectID()
	sensor, mqttClient, err = repo.ActivateExistingSensor(id, name, coordsX, coordsY, callback)
	assert.NoError(t, err)
	assert.NotNil(t, sensor)
	assert.Equal(t, id, sensor.Id)
	assert.NotNil(t, mqttClient)

	// Teste EmulateSensorData
}
