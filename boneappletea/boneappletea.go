package boneappletea

import (
	"fmt"
	"github.com/boneappletea/models"
	"github.com/boneappletea/mongo"
	"math/rand"
	"strings"
	"time"
)

//NOTE: call db here and return our struct
func Generate(sentence string) string {
	var words []string
	var newSentence = ""
	words = strings.Split(sentence, " ")

	for index, word := range words {
		//for _, word := range words {
		//NOTE: per db call, we create a new variable to use
		var bat models.Word

		//NOTE:query db for word
		bat = mongo.GetWord(strings.ToLower(word))

		//NOTE: if we get a hit out of the db, replace the word in the array
		if bat.Root != "" {
			//NOTE: our boneappletea is an rng of all the values supplied to the db
			length := len(bat.Values)
			//NOTE: if there are bats that have been approved
			if length > 0 {
				rng := rand.Intn(length)
				rand.Seed(time.Now().UnixNano())
				fmt.Println(length, rng) //NOTE: i was wondering if my rng was ranging right for the length of my array

				//NOTE:replace old sentence with new word/words
				words[index] = bat.Values[rng].Replacement
			}
		}
	}

	//NOTE: once we have the array, compress it back to a string
	newSentence = strings.Join(words, " ")

	return newSentence
}

//NOTE: get all bats in the db.  This should filter for words that have a flag of false.  That means they haven't been authenticated by us.
func Get() (int, []models.Word) {
	var code int
	var bats []models.Word
	bats = mongo.GetBats()

	if bats != nil {
		code = 200
	} else {
		code = 400
	}

	return code, bats
}

//NOTE: this is how a user adds a word to the database
func Add(bat models.Word) (int, string) {
	code, message := mongo.CreateBat(bat)

	return code, message
}

func Delete(word models.Word) (int, string) {
	code, message := mongo.DeleteBat(word)

	return code, message
}

func Accept(word models.Word) (int, string) {
	code, message := mongo.AcceptBat(word)

	return code, message
}
