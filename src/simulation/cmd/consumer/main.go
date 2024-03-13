package main

import (
	"context"
	"os"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/primary"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensorData"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
)

func init() {

	var variablesToCheck = [2]string{
		"RABBITMQ_URL",
		"MONGODB_URI",
	}

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
	} else {
		initialization.LoadEnvVariables(variablesToCheck[:])
	}

}

func main() {
	// go utils.MonitorNumberOfGoroutines()

	dbConnection, dbClient := infra.NewDBConnection()
	queueConnection := infra.NewQueueAdapter()
	queueConnection.GenerateConsumer("MQTTSensors")

	sensorDataAdapter := sensorData.NewSensorDataAdapter(dbConnection)
	messageHandlerAdapter := primary.NewMessageHandlerAdapter("MQTTSensors", queueConnection, sensorDataAdapter)

	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	messageHandlerAdapter.Consume()

}
