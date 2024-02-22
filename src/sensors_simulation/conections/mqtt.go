package conections

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GenerateClient(clientName string) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(clientName)
	opts.SetUsername(os.Getenv("RABBIT_USER"))
	opts.SetPassword(os.Getenv("RABBIT_PASSWORD"))

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Printf("->Client %s connected succecifuly!", clientName)
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %v", clientName, err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
