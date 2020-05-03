package api

import (
	"bytes"
	"fmt"
	"github.com/boneappletea/boneappletea"
	"github.com/boneappletea/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	//"math/rand"
	//"time"
)

func Start() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "POST", "GET"},
		//AllowHeaders: []string{"sentence", "root", "replacement"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

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
		word := models.Word{Root: strings.ToLower(c.GetHeader("root"))}
		value := models.Value{Flag: false, Replacement: strings.ToLower(c.GetHeader("value"))}
		word.Values = append(word.Values, value)
		fmt.Println(word)

		code, message := boneappletea.Add(word)

		c.JSON(code, gin.H{
			"message": message,
		})

	})

	//NOTE: after a word has been added, we still have to accept it via the web app
	router.POST("/accept", func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		fmt.Println(buf)

		c.JSON(200, gin.H{
			"boneappleteas": "/accept",
		})
	})

	//NOTE:removes a word (document) if the last entry is deleted, otherwise just removes the boneappletea entry in the array
	router.DELETE("/delete", func(c *gin.Context) {
		//var word models.Word
		//var values []string
		//var flag = false
		//values = append(values, strings.ToLower(c.GetHeader("value")))

		//word.Root = strings.ToLower(c.GetHeader("root"))
		//word.Values = values
		//word.Flag = flag

		//code, message := boneappletea.Delete(word)

		//c.JSON(code, gin.H{
		//	"message": message,
		//})
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
