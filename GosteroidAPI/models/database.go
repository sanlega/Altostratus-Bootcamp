package models

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/your_project_name/utils"
)

var DB *mongo.Database

func InitDatabase() {
    clientOptions := options.Client().ApplyURI(utils.GetEnv("MONGO_URI"))

    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database(utils.GetEnv("MONGO_DB"))
}
