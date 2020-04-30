package mongo

import (
	"context"
	"github.com/boneappletea/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func connect() *mongo.Client {
	//NOTE:top is dev, bottom is live
	//clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?connect=direct")
	mongodb_server := getEnv("MONGODB_SERVER", "mongo")
	mongodb_port := getEnv("MONGODB_PORT", "27017")
	mongodb_user := getEnv("MONGODB_USER", "root")
	mongodb_pass := getEnv("MONGODB_PASS", "example")

	mongodb_connect_uri := "mongodb://" + mongodb_user + ":" + mongodb_pass + "@" + mongodb_server + ":" + mongodb_port + "/?connect=direct"

	clientOpts := options.Client().ApplyURI(mongodb_connect_uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

//NOTE: get BATS that are flagged as false meaning that isn't been authenticated by us yet
func GetBats() []models.Word {
	var words []models.Word
	filter := bson.M{"flag": false}
	client := connect()

	collection := client.Database("boneappletea").Collection("words")
	cursor, err := collection.Find(context.TODO(), filter)

	err = cursor.All(context.TODO(), &words)

	if err != nil {
		log.Fatal(err)
	}

	client.Disconnect(context.TODO())

	return words
}

//NOTE: this is a private function.  Its job is to search the db for words that exist and return t/f
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

		_, err := collection.InsertOne(context.TODO(), word)
		client.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
			return 500, "boneappletea create failed"
		}
	} else {
		//NOTE: if word does exist, we need to update the document in the DB with a new value in the word in the values array
		code, result := updateBat(word)
		return code, result
	}

	return 200, "boneappletea created"
}

func updateBat(word models.Word) (int, string) {
	//NOTE: this is how we identify the document in the db we want to update
	filter := bson.M{"root": word.Root}

	/*NOTE : word.Values when we're creating no boneappleteas is always an array for returning purposes.
	However for new entries and updates, there is only ever 1 value in this array, in the array so thats why we use [0] element just to pull it out to push onto the db array field
	*/
	update := bson.D{
		{"$push", bson.D{
			{"values", word.Values[0]},
		}},
	}

	client := connect()
	collection := client.Database("boneappletea").Collection("words")

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return 500, "boneappletea update failed"
	}

	return 200, "boneappletea updated"
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func DeleteBat(word models.Word) (int, string) {
	found := checkForBat(word.Root)

	if found {
		// this is how we identify the document in the db we want to make a deletion on
		filter := bson.M{"root": word.Root}

		client := connect()
		collection := client.Database("boneappletea").Collection("words")

		batToDelete := word.Values[0]

		currentWord := GetWord(word.Root)
		batExists := contains(currentWord.Values, batToDelete)

		batsInDoc := len(currentWord.Values)

		if batsInDoc > 1 && batExists {
			// only remove the item provided
			update := bson.D{
				{"$pull", bson.D{
					{"values", batToDelete},
				}},
			}

			_, err := collection.UpdateOne(context.TODO(), filter, update)
			client.Disconnect(context.TODO())
			if err != nil {
				log.Fatal(err)
				return 500, "boneappletea delete failed"
			} else {
				return 200, "boneappletea deleted"
			}
		} else if batsInDoc == 1 && batExists {
			// otherwise we delete the document
			_, err := collection.DeleteOne(context.TODO(), filter)
			client.Disconnect(context.TODO())
			if err != nil {
				log.Fatal(err)
				return 500, "root delete failed"
			} else {
				return 200, "root deleted"
			}
		} else {
			client.Disconnect(context.TODO())
			return 500, "boneappletea not found"
		}
	} else {
		return 500, "root not found"
	}
}
