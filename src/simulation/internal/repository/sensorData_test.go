package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSensorDataSuccess(t *testing.T) {
	// Dados de entrada simulando o JSON recebido
	jsonData := []byte(`{"sensorId":"1234","data":[{"temperature":22.5,"humidity":60}],"createdAt":"2024-03-20T14:45:30Z"}`)

	// Chamando a função com os dados de teste
	result, err := NewSensorData(&jsonData)

	// Verificações
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1234", result.SensorId)
	// Adicione mais assertivas conforme necessário para verificar os campos de `result`
}

func TestNewSensorDataFailure(t *testing.T) {
	// Dados de entrada inválidos (JSON malformado)
	invalidJsonData := []byte(`{"sensorId":"1234","data":[{]}`)

	// Chamando a função com os dados de teste
	result, err := NewSensorData(&invalidJsonData)

	// Verificações
	assert.Error(t, err)
	assert.Nil(t, result)
}
