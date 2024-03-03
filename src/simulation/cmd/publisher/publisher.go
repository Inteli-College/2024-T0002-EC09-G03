package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/internal/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/database"
	sensorTable "github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/database/sensor"
	genSensors "github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/generator/sensors"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/sensors"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/publisher/utils"
)

var sensorsAmount int

func init() {

	switch {
	case len(os.Args) > 1:
		resul, err := strconv.Atoi(os.Args[1])
		if err != nil {
			initialization.LoadEnvVariables(&os.Args[1])
			sensorsAmount = 3
			return
		}
		sensorsAmount = resul
		initialization.LoadEnvVariables()

	case len(os.Args) > 2:
		resul, err := strconv.Atoi(os.Args[2])

		if err != nil {
			log.Fatal("Wrong amount of sensors")
		}
		sensorsAmount = resul
		initialization.LoadEnvVariables(&os.Args[1])

	default:
		initialization.LoadEnvVariables()
		sensorsAmount = 3

	}

}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	db := database.Connect()

	sensorTable := sensorTable.New(db)

	sensorsInstanciated := genSensors.GenerateSensors(sensorsAmount, &sensors.Sensors, sensorTable, &wg)

	log.Printf("Tempo: %f\n", time.Since(start).Seconds())
	log.Printf("Tamanho: %d\n", len(*sensorsInstanciated))

	for _, sensor := range *sensorsInstanciated {
		go sensor.Emulate()
	}

	go utils.MonitorNumberOfGoroutines()

	select {}
}
