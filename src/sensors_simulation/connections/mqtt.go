package connections

import (
	"fmt"
	"os"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mClient = sync.Mutex{}

func GenerateClient(clientName *string) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(*clientName)
	opts.SetUsername(os.Getenv("RABBIT_USER"))
	opts.SetPassword(os.Getenv("RABBIT_PASSWORD"))

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Printf("->Client %s connected succecifuly!\n", *clientName)
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %v\n", *clientName, err)
	}

	mClient.Lock()
	defer mClient.Unlock()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
