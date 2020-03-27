package mongo

import (
	"context"
	"github.com/boneappletea/models"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//func connect() *mongo.Client {
//}

func GetWord(root string) models.Word {
	fmt.Println(root)
	var word models.Word

	//clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?connect=direct")
	clientOpts := options.Client().ApplyURI("mongodb://tji1498a.com:27017/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOpts)

	filter := bson.M{"root": root}

	collection := client.Database("boneappletea").Collection("words")
	//err = collection.FindOne(context.TODO(), bson.M{"root": "degrees"}).Decode(&word)

	err = collection.FindOne(context.TODO(), filter).Decode(&word)
	if err != nil {
		log.Fatal(err)
	}

	//client.Disconnect(context.TODO())

	return word
}

func setup() {

}
