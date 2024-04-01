package infra

import (
	"fmt"
	"sync"
	"testing"
)

func TestMQTTLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var wg sync.WaitGroup
	clientCount := 10

	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			clientName := fmt.Sprintf("client%d", clientID)
			client := generateClient(&clientName, &wg)
			topic := "test/load"
			message := fmt.Sprintf("message from %s", clientName)
			token := client.Publish(topic, 0, false, message)
			token.Wait()
		}(i)
	}

	wg.Wait()
}
