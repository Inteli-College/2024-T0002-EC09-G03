package conections

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var broker_url string
var broker_port string
var rabbit_name string
var rabbit_password string

func configureClient(name string) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", os.Getenv("BROKER_URL"), os.Getenv("BROKER_PORT")))
	opts.SetClientID(name)
	opts.SetUsername(rabbit_name)     // Substitua pelo seu usuário
	opts.SetPassword(rabbit_password) // Substitua pela sua senha

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Printf("->Client %s connected succecifuly!", name)
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("-> Client %s disconnected due to: %v", name, err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro carregando o arquivo .env")
	}
	broker_url = os.Getenv("BROKER_URL")
	broker_port = os.Getenv("BROKER_PORT")
	rabbit_name = os.Getenv("RABBIT_USER")
	rabbit_password = os.Getenv("RABBIT_PASSWORD")

	if broker_port == "" || broker_url == "" || rabbit_name == "" || rabbit_password == "" {
		log.Fatal("Env variables not set.")
	}
}

func main() {
	client := configureClient("sensor")
	myData := struct {
		Data string `json:"data"`
	}{
		Data: "Hello",
	}
	myJSON, _ := json.Marshal(myData)

	for {
		client.Publish("sensor/data", 1, true, myJSON)
		time.Sleep(time.Duration(1 * time.Second))
	}
}
