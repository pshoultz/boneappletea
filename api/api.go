package api

import (
	"github.com/boneappletea/boneappletea"
	"github.com/boneappletea/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func Start() {
	router := gin.Default()

	//NOTE: test route for debugging ect...
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello from router",
		})
	})

	//NOTE: pass to mongo package to query words and the bat them up
	router.GET("/bat", func(c *gin.Context) {
		sentence := c.GetHeader("sentence")

		bat := boneappletea.Generate(sentence)

		c.JSON(200, gin.H{
			"boneappletea": bat,
		})
	})

	//NOTE: mongo call to db using mongo package pass values from gin.context
	router.POST("/add", func(c *gin.Context) {
		var word models.Word
		var values []string
		var flag = false
		values = append(values, strings.ToLower(c.GetHeader("value")))

		word.Root = strings.ToLower(c.GetHeader("root"))
		word.Values = values
		word.Flag = flag

		code, message := boneappletea.Add(word)

		c.JSON(code, gin.H{
			"message": message,
		})

	})

	//NOTE: remove a word or remove one of the values the array?
	router.DELETE("/delete", func(c *gin.Context) {
		//NOTE: this shouldn't delete the whole word but rather one of the indexs in the array
		/*
			the header should contain the root word and the value that is going to be deleted.  This word may already exist, we just want to remove one of the values in the array
			this is the c.GetHeader property i'm using above.  You can set this up with postman really easily.
		*/

		/*EXAMPLE:
		in the databse we have word called paul.  paul has 2 boneappletea's, we want to remove the [1] index
		root: "paul", values: ["abc", "qwe"], flag: true

		this function should remove "qwe" for example but keep "abc".  The whole document isn't gone, just the index of "qwe"

		NOTE:
		if this word only has one index and its deleted.  The whole word should be removed from the db.
		*/
	})

	//NOTE: we'll need a way to authorize words.  This will be hooked into our angular application where we login and get back words to authorize
	router.GET("/authorize", func(c *gin.Context) {
		//NOTE: for angular application will ask for all words whose flag property is set to false.
		//The angular application will let us mark these as true and the word can be used in the application.
	})

	router.Run(":8080")
}
