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

//DBinstance func
func DBinstance() *mongo.Client {
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    MongoDb := os.Getenv("MONGO_URI")

    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

    return client
}

var DBClient *mongo.Client = DBinstance()

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cluster := os.Getenv("MONGO_DB")
    var collection *mongo.Collection = client.Database(cluster).Collection(collectionName)

    return collection
}