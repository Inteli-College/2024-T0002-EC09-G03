package main

import (
	"os"

	// "github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/actors"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/actors"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/reader"
)

func init() {
	envPath := ""

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	initialization.LoadEnvVariables(envPath)
}

func main() {
	db := database.Connect()

	go actors.WatchNewSensorReq(actors.ChNewSensor)

	reader.Reader(db)

	select {}
}
