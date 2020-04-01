package boneappletea

import (
	"fmt"
	"github.com/boneappletea/models"
	"github.com/boneappletea/mongo"
	"math/rand"
	"strings"
)

//NOTE: call db here and return our struct
func Generate(sentence string) string {
	var words []string
	var newSentence = ""
	words = strings.Split(sentence, " ")

	for index, word := range words {
		//NOTE: per db call, we create a new variable to use
		var bat models.Word

		//NOTE:query db for word
		bat = mongo.GetWord(strings.ToLower(word))

		//NOTE: if we get a hit out of the db, replace the word in the array
		if bat.Root != "" {
			fmt.Println(bat)
			//NOTE: our boneappletea is an rng of all the values supplied to the db
			length := len(bat.Values)
			rng := rand.Intn(length)
			//fmt.Println(length, rng)//NOTE: i was wondering if my rng was ranging right for the length of my array

			//NOTE:replace old sentence with new word/words
			words[index] = bat.Values[rng]
		}
	}

	//NOTE: once we have the array, compress it back to a string
	newSentence = strings.Join(words, " ")

	return newSentence
}

func Add(bat models.Word) (int, string) {
	code, message := mongo.CreateBat(bat)

	return code, message
}

func Delete(word models.Word) (int, string) {
	code, message := mongo.DeleteBat(word)

	return code, message
}
