package mgdb

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDriver struct {
	client *mongo.Client
}

func NewMongoClient(clientDb any) *MongoDriver {
	if clientDb == nil {
		log.Fatal("failed to create mongo driver input is nil")
	}
	client, ok := clientDb.(*mongo.Client)
	if !ok {
		log.Fatalf("failed to create mongo driver %v", clientDb)
	}

	return &MongoDriver{
		client: client,
	}
}

func (db *MongoDriver) CreateCollection(dbName string, collectionName string) *mongo.Collection {
	return db.client.Database(dbName).Collection(collectionName)
}
