package boneappletea

import (
	"github.com/boneappletea/models"
	"github.com/boneappletea/mongo"
	"strings"
)

//NOTE: call db here and return our struct
func Generate(sentence string) models.Word {
	var words []string
	var bat models.Word
	words = strings.Split(sentence, " ")

	bat = mongo.GetWord(words[0])
	//for _, word := range words {
	//	bat = mongo.GetWord(word)
	//}

	return bat
}

//func Update(word models.Word) {
//
//}
//
//func Delete(word models.word) {
//
//}
