package sensors

import (
	"math/rand"
)

func SolarSimulation() ([]SensorData, error) {
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

	return []SensorData{
		{
			Measurament: valorRadiação,
			Unit:        "W/m^2",
			Material:    "light",
		},
	}, nil
}
