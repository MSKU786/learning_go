package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client{
    err := godotenv.Load()

    if (err != nil) {
      
      log.Fatal("Error loading the dotenv file", err)
    }

    MongoDB := os.Getenv("DB_HOST")

    client, err := mongo.NewClient(options.Client( ).ApplyURI(MongoDB))

    if (err != nil) {
      log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
    defer cancel()

    err = client.Connect(ctx)

    if (err != nil) {
      log.Fatal(err);
    }
    
    fmt.Println("Connected to MongoDB!")
    return client;
}


var Client *mongo.Client = DBInstance();

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
  return client.Database("cluster0").Collection(collectionName); 
}