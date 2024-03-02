package reader

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/consumer/database"
	"github.com/google/uuid"
)

func sendBatch(consumerControl chan struct{}, db dbInterface, wg *sync.WaitGroup, batch *[1000]*database.SensorsData, start *time.Time) {
	wg.Wait()

	slice := batch[:]
	db.CreateSensorsDataBatch(&slice)

	log.Printf("\033[35m[ %d ms ]\033[0m to group 1000 records and send to DB\n", time.Since(*start).Milliseconds())

	consumerControl <- struct{}{}
}

func feedBatch(msg *[]byte, batch *[1000]*database.SensorsData, index int, mArr *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	var sensorReceived Sensor
	err := json.Unmarshal(*msg, &sensorReceived)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err)
		return
	}
	dataToJson := struct {
		Data []SensorData `json:"data"`
	}{
		Data: sensorReceived.Data,
	}
	jsonData, _ := json.Marshal(dataToJson)
	sensorData := database.SensorsData{
		Id:        uuid.New().String(),
		Sensor_id: sensorReceived.Id,
		Data:      string(jsonData),
		CreatedAt: sensorReceived.Date,
	}

	mArr.Lock()
	batch[index] = &sensorData
	mArr.Unlock()
}
