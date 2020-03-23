package mongo

import (
	"context"
	"github.com/boneappletea/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//func connect() *mongo.Client {
//}

func GetWord(root string) models.Word {
	var words models.Word

	clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"root": "degrees"}
	collection := client.Database("boneappletea").Collection("words")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.FindOne(ctx, filter).Decode(&words)
	if err != nil {
		log.Fatal(err)
	}

	return words
}

func setup() {

}
