package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	mongoURI    = "mongodb://localhost:27017"
	MongoClient *mongo.Client
)

func InitializeMongodb() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println("Error initializing MongoDB client:", err)
		return
	}

	MongoClient = client

	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Println("error: failed to ping database")
		return
	}

	log.Println("connected to mongodb")

	checkDatabaseAndCollection()

}

func checkDatabaseAndCollection() {
	db := MongoClient.Database("employee_db")
	collectionName := "employees"

	if err := db.CreateCollection(context.Background(), collectionName); err != nil {
		if err.Error() != "namespace exists" {
			log.Println("Error creating collection:", err)
		}
	}
}
