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
)

func init() {
	envPath := ""

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	initialization.LoadEnvVariables(envPath)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	db := database.Connect()
	sensorsInstanciated := generator.GenerateSensors(1000, &sensors.Sensors, db, &wg)

	log.Printf("Tempo: %f\n", time.Since(start).Seconds())
	log.Printf("Tamanho: %d\n", len(*sensorsInstanciated))
	for _, sensor := range *sensorsInstanciated {
		// wgGroup := sync.WaitGroup{}
		//
		// go func(db *gorm.DB) {
		// 	wgGroup.Add(1)
		// 	database.CreateSensor(db, sensor.Name)
		// 	defer wgGroup.Done()
		// }(db)

		// wgGroup.Wait()
		go sensor.Emulate()
	}

	select {}
}
