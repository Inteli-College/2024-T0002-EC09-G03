package sensors

import (
	"math/rand"
	"time"
)

func GasSimulation() ([]SensorData, error) {
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

	return []SensorData{
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
