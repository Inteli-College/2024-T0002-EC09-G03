package reader

import (
	"log"
	"sync"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/consumer/database"
)

func Reader(db dbInterface, queue queueInterface) {
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
		start := time.Now()

		for i := 0; i < 1000; i++ {
			msg := queue.RetriveLastMessage()
			wg.Add(1)
			go feedBatch(msg, &batch, i, &mArr, &wg)

		}

		go sendBatch(consumerControl, db, &wg, &batch, &start)

	}

}
