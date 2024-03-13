package entity

import (
	"math/rand"
	"sync"
	"time"
)

func getTimePeriod() string {
	horaAtual := time.Now().Hour()
	switch {
	case horaAtual >= 6 && horaAtual < 12:
		return "Manhã"
	case horaAtual >= 12 && horaAtual < 18:
		return "Tarde"
	default:
		return "Noite"
	}
}

func PolutionSimulation() (*[]DataSensor, error) {
	// Gera valores aleatórios para PM2.5 e PM10
	// Os valores são baseados em uma faixa que pode ser considerada típica para ambientes urbanos
	// A faixa escolhida é completamente arbitrária e para fins de simulação apenas

	// Simula um aumento de partículas durante horários de pico de tráfego
	horaAtual := time.Now().Hour()

	var pm25, pm10 float64

	basePM25 := 5.0  // Base para PM2.5, representando um nível baixo/moderado
	basePM10 := 10.0 // Base para PM10, um pouco maior que PM2.5 devido a partículas maiores

	// Variações devido ao tráfego e outras fontes de partículas
	if (horaAtual >= 7 && horaAtual <= 9) || (horaAtual >= 17 && horaAtual <= 19) {
		pm25 = basePM25 + rand.Float64()*25 // Aumento durante o tráfego
		pm10 = basePM10 + rand.Float64()*50
	} else {
		pm25 = basePM25 + rand.Float64()*15 // Variação normal
		pm10 = basePM10 + rand.Float64()*30
	}

	return &[]DataSensor{
		{
			Measurament: pm25,
			Unit:        "µg/m³",
			Material:    "PM2.5",
		},
		{
			Measurament: pm10,
			Unit:        "µg/m³",
			Material:    "PM10",
		},
	}, nil
}

func GasSimulation() (*[]DataSensor, error) {
	time := time.Now().Hour()
	var co, nh3, no2 float64

	// CO: Maior entre 7-9h e 17-19h (horários de pico).
	if (time >= 7 && time <= 9) || (time >= 17 && time <= 19) {
		co = 300 + rand.Float64()*200 // 300 a 500 ppm durante horários de pico
	} else {
		co = 50 + rand.Float64()*250 // 50 a 300 ppm durante outros horários
	}

	// NH3: Varia aleatoriamente, representando fontes variadas.
	nh3 = 10 + rand.Float64()*40 // 10 a 50 ppm

	// NO2: Ligeiramente maior durante o dia devido à atividade humana e tráfego.
	if time >= 6 && time < 18 {
		no2 = 20 + rand.Float64()*80 // 20 a 100 ppm durante o dia
	} else {
		no2 = 10 + rand.Float64()*60 // 10 a 70 ppm durante a noite
	}

	return &[]DataSensor{
		{
			Measurament: co,
			Unit:        "ppm",
			Material:    "co",
		}, {
			Measurament: nh3,
			Unit:        "ppm",
			Material:    "nh3",
		}, {
			Measurament: no2,
			Unit:        "ppm",
			Material:    "no2",
		},
	}, nil
}

func SolarSimulation() (*[]DataSensor, error) {
	periodo := getTimePeriod()

	var valorRadiação float64

	switch periodo {
	case "Manhã":
		valorRadiação = rand.Float64() * 500
	case "Tarde":
		valorRadiação = 500 + rand.Float64()*500
	case "Noite":
		valorRadiação = rand.Float64() * 50
	}

	return &[]DataSensor{
		{
			Measurament: valorRadiação,
			Unit:        "W/m^2",
			Material:    "light",
		},
	}, nil
}

type SensorsTypes struct {
	Name     string
	Callback SensorCallbackFunc
}

type SensorClient struct {
	Sensor Sensor
	Client interface{}
}

type SensorsInstances []SensorClient

type Generator struct {
	SensorTypes    map[string]SensorsTypes
	MArr           *sync.Mutex
	Wg             *sync.WaitGroup
	Instances      *SensorsInstances
	ExistingCoords map[string]map[float64]bool
	MMap           sync.RWMutex
}

func NewGenerator() *Generator {

	sensorTypes := map[string]SensorsTypes{
		"Solar": {
			Name:     "Solar",
			Callback: SolarSimulation,
		},
		"Gas": {
			Name:     "Gas",
			Callback: GasSimulation,
		},
		"Polution": {
			Name:     "Polution",
			Callback: PolutionSimulation,
		},
	}

	return &Generator{
		SensorTypes:    sensorTypes,
		MArr:           &sync.Mutex{},
		Wg:             &sync.WaitGroup{},
		Instances:      &SensorsInstances{},
		ExistingCoords: make(map[string]map[float64]bool),
		MMap:           sync.RWMutex{},
	}
}
