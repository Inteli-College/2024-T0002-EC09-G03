package main

import (
	"os"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/connections/rabbitmq"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/reader"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/utils"
)

func init() {

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(&os.Args[1])
	} else {
		initialization.LoadEnvVariables()
	}

}

func main() {
	db := database.New()
	db.Connect()

	queue := rabbitmq.New()
	queue.GenerateConsumer("MQTTSensors")

	go utils.MonitorNumberOfGoroutines()

	reader.Reader(db, queue)

	select {}
}
