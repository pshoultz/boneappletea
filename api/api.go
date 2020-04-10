package api

import (
	"github.com/boneappletea/boneappletea"
	"github.com/boneappletea/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

func Start() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "POST", "GET"},
		AllowHeaders:  []string{"Origin"},
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
		var word models.Word
		var values []string
		var flag bool
		flag = false
		values = append(values, strings.ToLower(c.GetHeader("value")))

		word.Root = strings.ToLower(c.GetHeader("root"))
		word.Values = values
		word.Flag = flag
		word.ID = makeID()

		code, message := boneappletea.Add(word)

		c.JSON(code, gin.H{
			"message": message,
		})

	})

	//NOTE: after a word has been added, we still have to accept it via the web app
	router.POST("/accept", func(*gin.context) {
		/*NOTE:
		each word in the DB has an ID, we use the query from the post to get it.
		we use the ID here to find the document and then sent it flags to true.  Once true, the word is available for use in the mobile app
		*/
	})

	//NOTE:removes a word (document) if the last entry is deleted, otherwise just removes the boneappletea entry in the array
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

func makeID() string {
	rand.Seed(time.Now().UnixNano())
	length := 16
	bytes := make([]byte, length)
	characters := []rune("abcdefghijklmnopqrstuvwxyz123456789")
	for i := 0; i < length; i++ {
		bytes[i] = byte(characters[rand.Intn(35)])
	}

	return string(bytes)
}
