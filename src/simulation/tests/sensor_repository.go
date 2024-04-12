package tests

import (
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensor"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
)

func GenerateSensorAdapter() *sensor.SensorAdapter {
	dbConnection, _ := infra.NewDBConnection()

	return sensor.NewSensorDataAdapter(dbConnection)

}
