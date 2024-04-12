package repository

import (
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/tests"
)

func TestNewSensorRepo(t *testing.T) {
	path := "../../.env.test"

	initialization.LoadEnvVariables([]string{}, &path)

	t.Log("Generate support entities: dbConnection and SensorAdapter")

	sensorAdapter := tests.GenerateSensorAdapter()

	t.Log("Trying to create SensorRepo")

	sensorRepo := NewSensorRepo(sensorAdapter)

	if sensorRepo == nil {
		t.Fatal("Unable to create SensorRepo")
	}
}

func TestCreateSensor(t *testing.T) {
	path := "../../.env.test"

	initialization.LoadEnvVariables([]string{}, &path)

	t.Log("Generate support entities: dbConnection and SensorAdapter")

	sensorAdapter := tests.GenerateSensorAdapter()

	t.Log("Trying to create Sensor in db")

	sensorRepo := NewSensorRepo(sensorAdapter)

	sensor, mqttClient, err := sensorRepo.CreateSensor("Sensor-test", 99.99, 99.99, nil)

	if err != nil {
		t.Logf("Error creating sensor: %s\n", err.Error())
	}

	if sensor == nil || mqttClient == nil {
		t.Log("Not able to create sensor")
	}
}
