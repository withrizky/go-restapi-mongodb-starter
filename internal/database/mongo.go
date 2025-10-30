package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"myapp/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
	uri := config.GetEnv("MONGO_URI")
	dbName := config.GetEnv("DB_NAME")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ MongoDB connected successfully")
}
