package reader

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/connections/rabbitmq"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/google/uuid"
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

		var mArr = sync.Mutex{}
		var wg = sync.WaitGroup{}
		var batch [1000]*database.SensorsData

		for i := 0; i < 1000; i++ {
			msg, ok := <-msgs
			if !ok {
				log.Fatalln("Unable to retrive msg from queue.")
			}
			wg.Add(1)
			go feedBatch(&msg, &batch, i, &mArr, &wg)

		}

		go sendBatch(consumerControl, db, &wg, &batch)

	}

}

func sendBatch(consumerControl chan struct{}, db *gorm.DB, wg *sync.WaitGroup, batch *[1000]*database.SensorsData) {
	wg.Wait()

	slice := batch[:]
	database.CreateSensorsDataBatch(db, &slice)

	consumerControl <- struct{}{}
}

func feedBatch(msg *amqp091.Delivery, batch *[1000]*database.SensorsData, index int, mArr *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	var sensorReceived Sensor
	err := json.Unmarshal(msg.Body, &sensorReceived)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err)
		return
	}
	jsonData, _ := json.Marshal(sensorReceived.Data)
	sensorData := database.SensorsData{
		Id:           uuid.New().String(),
		Sensor_id:    sensorReceived.Id,
		Data:         string(jsonData),
		Coordinate_x: sensorReceived.CoordsX,
		Coordinate_y: sensorReceived.CoordsY,
		CreatedAt:    sensorReceived.Date,
	}

	mArr.Lock()
	batch[index] = &sensorData
	mArr.Unlock()
}
