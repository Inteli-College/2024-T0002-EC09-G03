package primary

import (
	"log"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/ports"
	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/repository"
)

type MessageHandlerAdapter struct {
	queue             string
	queueAdapter      ports.QueuePort
	sensorDataAdapter ports.SensorDataPort
}

func (s *MessageHandlerAdapter) Consume() {

	consumerControl := make(chan struct{}, 98)
	for i := 0; i < 98; i++ {
		consumerControl <- struct{}{}
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for {
		<-consumerControl

		var mArr = sync.Mutex{}
		var wg = sync.WaitGroup{}
		var batch [1000]*entity.SensorsData
		start := time.Now()

		for i := 0; i < 1000; i++ {
			msg := s.queueAdapter.RetriveLastMessage(s.queue)
			wg.Add(1)
			go feedBatch(msg, &batch, i, &mArr, &wg)

		}

		go sendBatch(consumerControl, s.sensorDataAdapter, &wg, &batch, &start)

	}
}

func sendBatch(consumerControl chan struct{}, db ports.SensorDataPort, wg *sync.WaitGroup, batch *[1000]*entity.SensorsData, start *time.Time) {
	wg.Wait()

	slice := batch[:]
	err := db.CreateInBatch(&slice)

	if err != nil {
		log.Fatalf("Unable to send batch: %s", err.Error())
	}

	log.Printf("\033[35m[ %d ms ]\033[0m to group 1000 records and send to DB\n", time.Since(*start).Milliseconds())

	consumerControl <- struct{}{}
}

func feedBatch(msg *[]byte, batch *[1000]*entity.SensorsData, index int, mArr *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	sensorData, err := repository.NewSensorData(msg)

	if err != nil {
		log.Println(err)
		return
	}

	mArr.Lock()
	batch[index] = sensorData
	mArr.Unlock()
}

func NewMessageHandlerAdapter(queue string, queueAdapter ports.QueuePort, sensorDataAdapter ports.SensorDataPort) *MessageHandlerAdapter {
	return &MessageHandlerAdapter{
		queue:             queue,
		queueAdapter:      queueAdapter,
		sensorDataAdapter: sensorDataAdapter,
	}
}
