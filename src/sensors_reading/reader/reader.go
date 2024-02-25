package reader

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/connections/rabbitmq"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Sensor struct {
	Name    string       `json:"name"`
	Id      string       `json:"id"`
	Data    []SensorData `json:"data"`
	CoordsX float64      `json:"coords_x"`
	CoordsY float64      `json:"coords_y"`
	Date    time.Time    `json:"date"`
}

type SensorData struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func Reader(db *gorm.DB) {
	msgs := rabbitmq.GenerateConsumer("MQTTSensors")
	if msgs == nil {
		log.Fatal("Msgs null")
	}

	consumerControl := make(chan struct{}, 98)
	for i := 0; i < 98; i++ {
		consumerControl <- struct{}{}
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for {
		<-consumerControl

		msg, ok := <-msgs

		if !ok {
			log.Fatalln("Unable to retrive msg from queue.")
		}
		go func(msg *amqp091.Delivery, db *gorm.DB) {
			defer func() { consumerControl <- struct{}{} }()
			var sensorReceived Sensor
			err := json.Unmarshal(msg.Body, &sensorReceived)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				return
			}

			dataJson, _ := json.Marshal(sensorReceived.Data)

			database.CreateSensorsData(db, sensorReceived.Id, string(dataJson), sensorReceived.CoordsX, sensorReceived.CoordsY, sensorReceived.Date)

		}(&msg, db)
	}

}
