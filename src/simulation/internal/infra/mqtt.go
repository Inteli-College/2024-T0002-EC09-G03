package infra

import (
	"fmt"
	"os"
	"sync"

	mqttPaho "github.com/eclipse/paho.mqtt.golang"
)

func generateClient(clientName *string, wg *sync.WaitGroup) mqttPaho.Client {
	opts := mqttPaho.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(*clientName)
	opts.SetUsername(os.Getenv("RABBIT_USER"))
	opts.SetPassword(os.Getenv("RABBIT_PASSWORD"))

	opts.OnConnect = func(client mqttPaho.Client) {
		fmt.Printf("-> Client %s connected successfully!\n", *clientName)
		defer wg.Done()
	}

	opts.OnConnectionLost = func(client mqttPaho.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %s\n", *clientName, err.Error())
		wg.Add(1)
	}

	client := mqttPaho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Erro na conexão: %s\n", token.Error()))
	}

	wg.Wait()

	return client
}

func NewMQTTConnetion(clientName *string, wg *sync.WaitGroup) mqttPaho.Client {
	wg.Add(1)
	return generateClient(clientName, wg)
}
