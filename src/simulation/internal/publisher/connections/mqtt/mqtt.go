package mqtt

import (
	"fmt"
	"log"
	"os"
	"sync"

	mqttPaho "github.com/eclipse/paho.mqtt.golang"
)

var mClient = sync.Mutex{}

type mqtt struct {
	clientName *string
	client     mqttPaho.Client
}

func (s *mqtt) generateClient(wg *sync.WaitGroup) {
	opts := mqttPaho.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(*s.clientName)
	opts.SetUsername(os.Getenv("RABBIT_USER"))
	opts.SetPassword(os.Getenv("RABBIT_PASSWORD"))

	opts.OnConnect = func(client mqttPaho.Client) {
		fmt.Printf("-> Client %s connected successfully!\n", *s.clientName)
		defer wg.Done()
	}

	opts.OnConnectionLost = func(client mqttPaho.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %s\n", *s.clientName, err.Error())
		wg.Add(1)
	}

	// mClient.Lock()
	// defer mClient.Unlock()
	client := mqttPaho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Erro na conexão: %s\n", token.Error()))
	}

	s.client = client
}

func (s *mqtt) Publish(topic *string, qos byte, retained bool, payload interface{}) {
	if token := s.client.Publish(*topic, qos, retained, payload); token.Wait() && token.Error() != nil {

		log.Fatalf("From: %s - error: %s", *s.clientName, token.Error().Error())
	}
}

func New(clientName *string, wg *sync.WaitGroup) *mqtt {
	wg.Add(1)
	mqtt := mqtt{
		clientName: clientName,
	}

	mqtt.generateClient(wg)

	return &mqtt
}
