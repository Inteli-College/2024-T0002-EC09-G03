package infra

import (
	"context"
	"fmt"
	"testing"

	initialization "github.com/Inteli-College/2024-T0002-EC09-G03/init"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDBLoad(t *testing.T) {
	fmt.Printf("TESTANDO O LOAD DB")

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	path := "./.env"
	initialization.LoadEnvVariables([]string{"MONGODB_URI"}, &path)

	db, client := NewDBConnection()
	defer client.Disconnect(context.TODO())

	collection := db.Collection("testCollection")

	for i := 0; i < 10; i++ {
		doc := bson.D{{Key: "key", Value: fmt.Sprintf("value%d", i)}}
		if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
			t.Fatalf("Failed to insert document: %v", err)
		}
	}
	fmt.Printf("FIM DO TESTANDO O LOAD DB")
}
