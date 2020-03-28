package api

import (
	"fmt"
	"github.com/boneappletea/boneappletea"
	"github.com/boneappletea/models"
	"github.com/gin-gonic/gin"
	"strconv"
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
		var flag, _ = strconv.ParseBool(c.GetHeader("flag"))
		var likes = 0
		var dislikes = 0

		values = append(values, strings.ToLower(c.GetHeader("value")))
		word.Root = strings.ToLower(c.GetHeader("root"))
		word.Values = values
		word.Flag = flag
		word.Likes = likes
		word.Dislikes = dislikes

		fmt.Println(word)

		code, message := boneappletea.Add(word)

		c.JSON(code, gin.H{
			"message": message,
		})

	})

	//NOTE: remove a word or remove one of the values the array?
	router.DELETE("/delete", func(c *gin.Context) {

	})

	router.Run(":8080")
}
