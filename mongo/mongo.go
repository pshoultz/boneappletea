package mongo

import (
	"context"
	"github.com/boneappletea/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func connect() *mongo.Client {
	//clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?connect=direct")
	clientOpts := options.Client().ApplyURI("mongodb://tji1498a.com:27017/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetWord(root string) models.Word {
	var word models.Word
	filter := bson.M{"root": root}

	//clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?connect=direct")
	//clientOpts := options.Client().ApplyURI("mongodb://tji1498a.com:27017/?connect=direct")
	//client, err := mongo.Connect(context.TODO(), clientOpts)
	client := connect()

	collection := client.Database("boneappletea").Collection("words")

	collection.FindOne(context.TODO(), filter).Decode(&word)

	client.Disconnect(context.TODO())

	return word
}
