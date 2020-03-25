package api

import (
	"github.com/boneappletea/boneappletea"
	"github.com/boneappletea/models"
	"github.com/gin-gonic/gin"
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
		values = append(values, c.Query("value"))

		word.Root = c.Query("root")
		word.Values = values
		word.Flag = false
		word.Likes = 0
		word.Dislikes = 0
	})

	//NOTE: remove a word or remove one of the values the array?
	router.DELETE("/delete", func(c *gin.Context) {

	})

	router.Run(":8080")
}
