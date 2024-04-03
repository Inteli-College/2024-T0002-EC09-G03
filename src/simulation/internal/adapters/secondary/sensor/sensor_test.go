package sensor

import (
	"context"
	"os"
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"

	"github.com/Inteli-College/2024-T0002-EC09-G03/internal/domain/entity"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateSensor(t *testing.T) {

	path := "./.env"
	initialization.LoadEnvVariables([]string{"MONGODB_URI"}, &path)
	uri := os.Getenv("MONGOBD_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	require.NoError(t, err)
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	coll := db.Collection("testdb")

	// Limpa a coleção após o teste
	defer coll.Drop(context.Background())

	adapter := NewSensorDataAdapter(db)

	// Criação de um sensor para teste
	testSensor := &entity.Sensor{
		Name:        "Test Sensor",
		Id:          primitive.NilObjectID,
		CoordinateX: 2,
		CoordinateY: 3,
	}

	createdSensor, err := adapter.CreateSensor(testSensor)
	require.NoError(t, err)
	require.NotNil(t, createdSensor)
	require.NotEqual(t, primitive.NilObjectID, createdSensor.Id)
}

func TestGetAllSensors(t *testing.T) {

	path := "./.env"
	initialization.LoadEnvVariables([]string{"MONGODB_URI"}, &path)
	uri := os.Getenv("MONGOBD_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	require.NoError(t, err)
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	coll := db.Collection("testdb")

	// Limpa a coleção após o teste
	defer coll.Drop(context.Background())

	adapter := NewSensorDataAdapter(db)

	// Inserindo dados de teste diretamente para validar o GetAllSensors
	testSensor := entity.Sensor{
		Name:        "Test Sensor",
		Id:          primitive.NilObjectID,
		CoordinateX: 2,
		CoordinateY: 3,
	}
	_, err = coll.InsertOne(context.Background(), testSensor)
	require.NoError(t, err)

	sensors := adapter.GetAllSensors()
	require.NotNil(t, sensors)
	require.NotEmpty(t, sensors)
	require.Equal(t, "Test Sensor", (*sensors)[0].Name)
}
