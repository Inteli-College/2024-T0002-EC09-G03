package sensors

import "math/rand"

func SolarSimulation() (float64, error) {
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

	return valorRadiação, nil
}
