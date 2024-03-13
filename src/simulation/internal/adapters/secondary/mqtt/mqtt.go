package mqtt

import (
	"log"
	"sync"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/infra"
	mqttPaho "github.com/eclipse/paho.mqtt.golang"
)

type MQTTAdapter struct {
	clientName *string
	client     mqttPaho.Client
}

func (s *MQTTAdapter) CreateClient(clientName *string) {
	s.clientName = clientName

	if s.client != nil {
		panic("Trying to overide existing client")
	}

	wg := sync.WaitGroup{}
	s.client = infra.NewMQTTConnetion(clientName, &wg)
}

func (s *MQTTAdapter) Publish(topic *string, qos byte, retained bool, payload interface{}) {
	if s.client == nil {
		panic("MQTT client not initialized!")
	}

	if token := s.client.Publish(*topic, qos, retained, payload); token.Wait() && token.Error() != nil {

		log.Fatalf("From: %s - error: %s", *s.clientName, token.Error().Error())
	}
}
