package mongo

import (
	"context"
	"fmt"
	"github.com/boneappletea/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func connect() {
	var (
		client   *mongo.Client
		mongoURL = "mongodb://localhost:27017"
	)

	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}

	fmt.Println("connected to nosql database:", mongoURL)
}

func GetWord(root string) models.Word {
	var word models.Word
	connect()
	return word
}

func setup() {

}
