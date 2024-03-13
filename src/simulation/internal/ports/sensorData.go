package ports

import "github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"

type SensorDataPort interface {
	CreateInBatch(*[]*entity.SensorsData) error
}
