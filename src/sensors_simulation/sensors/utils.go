package sensors

import "time"

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
