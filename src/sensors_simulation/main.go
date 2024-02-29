package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/database"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/generator"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/sensors"
	"github.com/Inteli-College/2024-T0002-EC09-G03/src/sensors_simulation/utils"
)

func init() {

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(&os.Args[1])
	} else {
		initialization.LoadEnvVariables()
	}

}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	db := database.Connect()
	sensorsInstanciated := generator.GenerateSensors(3, &sensors.Sensors, db, &wg)

	log.Printf("Tempo: %f\n", time.Since(start).Seconds())
	log.Printf("Tamanho: %d\n", len(*sensorsInstanciated))

	for _, sensor := range *sensorsInstanciated {
		go sensor.Emulate()
	}

	go utils.MonitorNumberOfGoroutines()

	select {}
}
