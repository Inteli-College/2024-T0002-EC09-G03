package actors

import (
	"log"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"gorm.io/gorm"
)

type ChannelMessage struct {
  Name string
	DB    *gorm.DB
	ChOut chan string
}

var ChNewSensor = make(chan ChannelMessage)

func WatchNewSensorReq(incomingCh chan ChannelMessage) {
	for {
		var sensorId string
		sensor, ok := <-incomingCh

		if !ok {
			log.Println("Channel was closed!")
			break
		}

		if check := database.GetSensorByName(sensor.DB, sensor.Name); check.Name == "" {
			createdSensor := database.CreateSensor(sensor.DB, sensor.Name)
			sensorId = createdSensor.Id
		} else {
			sensorId = check.Id
		}

		sensor.ChOut <- sensorId

	}
}
