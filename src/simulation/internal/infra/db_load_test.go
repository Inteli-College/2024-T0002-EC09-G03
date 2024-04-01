package infra

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestDBLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db, client := NewDBConnection()
	defer client.Disconnect(context.TODO())

	collection := db.Collection("testCollection")

	for i := 0; i < 100; i++ {
		doc := bson.D{{Key: "key", Value: fmt.Sprintf("value%d", i)}}
		if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
			t.Fatalf("Failed to insert document: %v", err)
		}
	}
}
