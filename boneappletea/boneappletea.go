package boneappletea

import (
	"fmt"
	"github.com/boneappletea/models"
	"github.com/boneappletea/mongo"
	"strings"
)

//NOTE: call db here and return our struct
func Generate(sentence string) string {
	var words []string
	var newSentence = ""
	words = strings.Split(sentence, " ")

	for _, word := range words {
		var bat models.Word
		bat = mongo.GetWord(strings.ToLower(word))
		fmt.Println(word, bat)
	}

	//NOTE: 2nd for loop to find the word to replace from original sentence

	return newSentence
}

//func Update(word models.Word) {
//
//}
//
//func Delete(word models.word) {
//
//}
