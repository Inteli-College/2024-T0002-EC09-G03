package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensor"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/ports"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/repository"
	"github.com/joho/godotenv"
)

var sensorsAmount int = 3

func init() {

  var variablesToCheck = [9]string{
    "BROKER_URL",
    "BROKER_PORT",
    "RABBIT_USER",
    "RABBIT_PASSWORD",
		"DATABASE_HOST",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
		"DATABASE_PORT",
	}

  switch {
  case len(os.Args) == 2:
		initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
  case len(os.Args) > 2:
    initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
    resul, err := strconv.Atoi(os.Args[2])
  default:
    initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
  
}

	if len(os.Args) > 1 {
		initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
	} else {
		initialization.LoadEnvVariables(variablesToCheck[:])
	}

	godotenv.Load(os.Args[1])

	// switch {
	// case len(os.Args) > 1:
	// 	resul, err := strconv.Atoi(os.Args[1])
	// 	if err != nil {
	// 		initialization.LoadEnvVariables(&os.Args[1])
	// 		sensorsAmount = 3
	// 		return
	// 	}
	// 	sensorsAmount = resul
	// 	initialization.LoadEnvVariables()
	//
	// case len(os.Args) > 2:
	// 	resul, err := strconv.Atoi(os.Args[2])
	//
	// 	if err != nil {
	// 		log.Fatal("Wrong amount of sensors")
	// 	}
	// 	sensorsAmount = resul
	// 	initialization.LoadEnvVariables(&os.Args[1])
	//
	// default:
	// 	initialization.LoadEnvVariables()
	// 	sensorsAmount = 3
	//
	// }

}

func main() {

	dbConnection := infra.NewDBConnection()

	sensorAdapter := sensor.NewSensorDataAdapter(dbConnection)

	generatorEntity := entity.NewGenerator()

	sensorRepo := repository.NewSensorRepo(sensorAdapter)

	generatorRepo := repository.NewGeneratorRepo(sensorAdapter, generatorEntity, sensorRepo)

	sensorsInstances := generatorRepo.GenerateSensors(3)

	for _, sensorClient := range *sensorsInstances {
		go func(sensorClient entity.SensorClient) {
			log.Printf("Emulando sensor: %s\n", sensorClient.Sensor.Name)

			if _, ok := sensorClient.Client.(ports.MQTTPort); !ok {
				panic(fmt.Sprintf("Client of %s should be implementation of ports.MQTTPort", sensorClient.Sensor.Id))
			}

			go sensorRepo.EmulateSensor(&sensorClient.Sensor, sensorClient.Client)
		}(sensorClient)
	}

	select {}
}
