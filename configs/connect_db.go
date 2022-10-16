package configs

import (
	"context"
	"log"
	"time"

	// "log"
	// "os"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	// err := godotenv.Load("fly.toml")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// dbUserName := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// urlDB := "mongodb+srv://" + dbUserName + ":" + dbPassword + "@cluster0.h98mko5.mongodb.net/?retryWrites=true&w=majority"
	urlDB := "mongodb+srv://golang_mongodb:wxXPUaajy9mAYB2W@cluster0.h98mko5.mongodb.net/?retryWrites=true&w=majority?ssl=true"
	client, err := mongo.NewClient(options.Client().ApplyURI(urlDB))
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("UserProfile").Collection(collectionName)
	return collection
}
