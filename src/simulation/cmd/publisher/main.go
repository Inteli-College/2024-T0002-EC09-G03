package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/adapters/secondary/sensor"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/ports"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/repository"
)

var sensorsAmount int = 3

func init() {

	argsSize := len(os.Args)

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

	switch argsSize {
	case 2:
		// Nao conseguiu converter para int
		if resul, err := strconv.Atoi(os.Args[1]); err != nil {
			initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
		} else {
			sensorsAmount = resul
			initialization.LoadEnvVariables(variablesToCheck[:])
		}

	case 3:
		initialization.LoadEnvVariables(variablesToCheck[:], &os.Args[1])
		if resul, err := strconv.Atoi(os.Args[2]); err == nil {
			sensorsAmount = resul
		}

	default:
		initialization.LoadEnvVariables(variablesToCheck[:])
	}

}

func main() {

	dbConnection := infra.NewDBConnection()

	sensorAdapter := sensor.NewSensorDataAdapter(dbConnection)

	generatorEntity := entity.NewGenerator()

	sensorRepo := repository.NewSensorRepo(sensorAdapter)

	generatorRepo := repository.NewGeneratorRepo(sensorAdapter, generatorEntity, sensorRepo)

	sensorsInstances := generatorRepo.GenerateSensors(sensorsAmount)

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
