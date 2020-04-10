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
		var flag bool
		flag = false
		values = append(values, strings.ToLower(c.GetHeader("value")))

		word.Root = strings.ToLower(c.GetHeader("root"))
		word.Values = values
		word.Flag = flag

		code, message := boneappletea.Add(word)

		c.JSON(code, gin.H{
			"message": message,
		})

	})

	// removes a word (document) if the last entry is deleted, otherwise just removes the boneappletea entry in the array
	router.DELETE("/delete", func(c *gin.Context) {
		var word models.Word
		var values []string
		var flag = false
		values = append(values, strings.ToLower(c.GetHeader("value")))

		word.Root = strings.ToLower(c.GetHeader("root"))
		word.Values = values
		word.Flag = flag

		code, message := boneappletea.Delete(word)

		c.JSON(code, gin.H{
			"message": message,
		})
	})

	//NOTE: we'll need a way to authorize words.  This will be hooked into our angular application where we login and get back words to authorize
	router.GET("/authorize", func(c *gin.Context) {
		//NOTE: for angular application will ask for all words whose flag property is set to false.
		//The angular application will let us mark these as true and the word can be used in the application.
	})

	//NOTE: gets all boneappleteas in the db who have flags set to false
	router.GET("/get", func(c *gin.Context) {
		var words []models.Word
		var code int

		code, words = boneappletea.Get()

		c.JSON(code, gin.H{
			"boneappleteas": words,
		})
	})

	router.Run(":8080")
}
