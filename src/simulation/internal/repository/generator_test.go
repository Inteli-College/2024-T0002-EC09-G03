package repository

import (
	"testing"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Definindo um mock para SensorPort
type MockSensorPort struct {
	mock.Mock
}

func (m *MockSensorPort) GetAllSensors() *[]entity.Sensor {
	args := m.Called()
	return args.Get(0).(*[]entity.Sensor)
}

func (m *MockSensorPort) CreateSensor(sensor *entity.Sensor) (*entity.Sensor, error) {
	args := m.Called(sensor)
	return args.Get(0).(*entity.Sensor), args.Error(1)
}

// Teste básico para GenerateSensors
func TestGenerateSensors(t *testing.T) {
	mockSensorPort := new(MockSensorPort)
	mockSensors := make([]entity.Sensor, 5) // Considerando que existam 5 sensores disponíveis
	mockSensorPort.On("GetAllSensors").Return(&mockSensors)

	sensorRepo := &SensorRepo{} // Isto pode precisar ser mockado ou configurado adicionalmente
	generatorRepo := GeneratorRepo{
		sensorAdapter: mockSensorPort,
		sensorRepo:    sensorRepo,
		gen:           &entity.Generator{},
	}

	numberToGenerate := 3
	generatedSensors := generatorRepo.GenerateSensors(numberToGenerate)
	assert.NotNil(t, generatedSensors)
	assert.Equal(t, numberToGenerate, len(*generatedSensors), "Deve gerar o número correto de sensores")
}

// Testando entradas inválidas para GenerateSensors
func TestGenerateSensorsInvalidInput(t *testing.T) {
	mockSensorPort := new(MockSensorPort)
	mockSensorPort.On("GetAllSensors").Return(&[]entity.Sensor{})

	sensorRepo := &SensorRepo{} // Isso pode precisar ser mockado ou configurado adicionalmente
	generatorRepo := GeneratorRepo{
		sensorAdapter: mockSensorPort,
		sensorRepo:    sensorRepo,
		gen:           &entity.Generator{},
	}

	testCases := []struct {
		name             string
		numberToGenerate int
		expectError      bool
	}{
		{"NegativeNumber", -1, true},
		{"ZeroNumber", 0, true},
		{"ExcessiveNumber", 1000, true}, // Assumindo 1000 como um número irrealisticamente alto
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			generatedSensors := generatorRepo.GenerateSensors(tc.numberToGenerate)
			if tc.expectError {
				assert.Nil(t, generatedSensors, "Deve retornar nil para entradas inválidas")
			} else {
				assert.NotNil(t, generatedSensors, "Deve gerar sensores para entradas válidas")
			}
		})
	}
}

// Testando estados inesperados das dependências
func TestGenerateSensorsUnexpectedDependencyState(t *testing.T) {
	mockSensorPort := new(MockSensorPort)
	mockSensorPort.On("GetAllSensors").Return(nil) // Simulando uma falha ou estado inesperado

	sensorRepo := &SensorRepo{} // Isso pode precisar ser mockado ou configurado adicionalmente
	generatorRepo := GeneratorRepo{
		sensorAdapter: mockSensorPort,
		sensorRepo:    sensorRepo,
		gen:           &entity.Generator{},
	}

	generatedSensors := generatorRepo.GenerateSensors(5)
	assert.Nil(t, generatedSensors, "Deve retornar nil quando as dependências estão em um estado inesperado")
}

// Testando a integração entre componentes
func TestIntegrationBetweenComponents(t *testing.T) {
	mockSensorPort := new(MockSensorPort)
	mockSensors := make([]entity.Sensor, 10) // Simulando 10 sensores disponíveis
	mockSensorPort.On("GetAllSensors").Return(&mockSensors)

	sensorRepo := &SensorRepo{} // Este mock pode ser expandido com comportamentos específicos se necessário
	generatorRepo := GeneratorRepo{
		sensorAdapter: mockSensorPort,
		sensorRepo:    sensorRepo,
		gen:           &entity.Generator{},
	}

	// Assumindo que a lógica de negócio permite gerar até 5 sensores sem problemas
	generatedSensors := generatorRepo.GenerateSensors(5)
	assert.NotNil(t, generatedSensors, "Deve gerar sensores quando as condições são ideais")
	assert.Equal(t, 5, len(*generatedSensors), "Deve respeitar a quantidade solicitada de sensores")
}
