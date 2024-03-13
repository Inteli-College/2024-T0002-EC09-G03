package main

import (
	"os"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/primary"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensorData"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
)

func init() {

	var variablesToCheck = [6]string{
		"RABBITMQ_URL",
		"DATABASE_HOST",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
		"DATABASE_PORT",
	}

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
	} else {
		initialization.LoadEnvVariables(variablesToCheck[:])
	}

}

func main() {
	// go utils.MonitorNumberOfGoroutines()

	dbConnection := infra.NewDBConnection()
	queueConnection := infra.NewQueueAdapter()
	queueConnection.GenerateConsumer("MQTTSensors")

	sensorDataAdapter := sensorData.NewSensorDataAdapter(dbConnection)
	messageHandlerAdapter := primary.NewMessageHandlerAdapter("MQTTSensors", queueConnection, sensorDataAdapter)

	messageHandlerAdapter.Consume()
}
