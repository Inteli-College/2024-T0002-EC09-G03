package utils

import (
	"log"
	"runtime"
	"time"
)

func MonitorNumberOfGoroutines() {
	for {
		time.Sleep(time.Second * 10)
		log.Printf("\033[33mGoroutines running: %d\033[0m\n", runtime.NumGoroutine())
		log.Printf("\033[33mNum of CPUs: %d\033[0m\n", runtime.NumCPU())
	}
}
