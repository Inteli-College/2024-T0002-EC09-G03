package sensorData

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateInBatch(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("yourDatabaseName")
	adapter := NewSensorDataAdapter(db)

	// Criando dados de teste para inserir
	testData := []*entity.SensorsData{
		{
			SensorId: primitive.NewObjectID(),
			Data: []entity.DataSensor{
				{
					Measurament: 23.5,
					Unit:        "Celsius",
					Material:    "Air",
				},
				{
					Measurament: 75.0,
					Unit:        "%",
					Material:    "Humidity",
				},
			},
			CreatedAt: time.Now(),
		},
	}

	err = adapter.CreateInBatch(&testData)
	if err != nil {
		t.Errorf("Failed to create data in batch: %v", err)
	}

	// Adicione aqui sua lógica de verificação para confirmar que os dados foram inseridos corretamente.
	// Verificação: Confirme que os dados foram inseridos corretamente
	for _, data := range testData {
		var result entity.SensorsData
		filter := bson.M{"sensor_id": data.SensorId}
		err := db.Collection("yourCollectionName").FindOne(context.TODO(), filter).Decode(&result)
		require.NoError(t, err, "Should be able to find the inserted data")
		assert.NotEqual(t, primitive.NilObjectID, result.Id, "The result should have a valid MongoDB ID")
		assert.Equal(t, data.SensorId, result.SensorId, "The SensorId should match")
		// Adicione mais verificações conforme necessário
	}

	// Limpeza: Remova os dados de teste
	for _, data := range testData {
		// Assume que você está usando o SensorId como filtro para identificar os dados inseridos
		filter := bson.M{"sensor_id": data.SensorId}
		_, err := db.Collection("yourCollectionName").DeleteMany(context.TODO(), filter)
		require.NoError(t, err, "Should be able to clean up the test data")
	}
}
