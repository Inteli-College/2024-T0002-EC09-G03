package main

import (
	"os"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/primary"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensorData"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/internal/init"
)

func init() {

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(&os.Args[1])
	} else {
		initialization.LoadEnvVariables()
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
