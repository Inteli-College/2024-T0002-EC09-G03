package sensors

import (
	"math/rand"
	"time"
)

func PolutionSimulation() ([]SensorData, error) {
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

	return []SensorData{
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
