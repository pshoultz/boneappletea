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

func checkForBat(root string) bool {
	var word models.Word
	filter := bson.M{"root": root}
	client := connect()

	collection := client.Database("boneappletea").Collection("words")
	collection.FindOne(context.TODO(), filter).Decode(&word)

	client.Disconnect(context.TODO())

	if word.Root == "" {
		return false
	}

	return true
}

func GetWord(root string) models.Word {
	var word models.Word
	filter := bson.M{"root": root}

	client := connect()

	collection := client.Database("boneappletea").Collection("words")

	collection.FindOne(context.TODO(), filter).Decode(&word)

	client.Disconnect(context.TODO())

	return word
}

func CreateBat(word models.Word) (int, string) {
	found := checkForBat(word.Root)

	//NOTE:: if word doesn't exist, we add it to the db
	if !found {
		client := connect()
		collection := client.Database("boneappletea").Collection("words")

		//result, err := collection.InsertOne(context.TODO(), word)
		//fmt.Println(result)
		_, err := collection.InsertOne(context.TODO(), word)

		client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
			return 500, "fail"
		}
	} else {
		//NOTE: if word does exist, we need to update the document in the DB with a new value in the word in the values array
		code, result := updateBat(word)
		return code, result
	}

	return 200, "success"
}

func updateBat(word models.Word) (int, string) {
	filter := bson.M{"root": word.Root}

	/*NOTE : word.Values when we're creating no boneappleteas is always an array for returning purposes.
	However for new entries and updates, there is only ever 1 variable, in the array so thats why we use [0] element just to pull it out to push onto the db array field
	*/
	update := bson.D{
		{"$push", bson.D{
			{"values", word.Values[0]},
		}},
	}

	client := connect()
	collection := client.Database("boneappletea").Collection("words")

	//result, err := collection.InsertOne(context.TODO(), filter, update)
	//fmt.Println(result)
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return 500, "fail"
	}

	return 200, "success"
}
