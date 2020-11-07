package main

import (
	"context"
    "fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func setupRouter() *gin.Engine {
	// Creates a router without any middleware by default
	// r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	// authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// 	authorized.POST("/read", readEndpoint)

	// 	// nested group
	// 	testing := authorized.Group("testing")
	// 	testing.GET("/analytics", analyticsEndpoint)
	// }

	// Default With the Logger and Recovery middleware already attached
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func setupDB()*mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("pttrainer")
	fmt.Println("Connected to MongoDB!")
	return database
}


func main() {
	setupDB()
	r := setupRouter()

	// r.GET("/clients", func(c *gin.Context) {
	// 	collection := db.Collection("qux")
	// 	opts := options.Find()
	// 	cursor, err := collection.Find(context.TODO(), bson.D).Decode(&post)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	var results []bson.M
	// 	if err = cursor.All(context.TODO(), &results); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, result := range results {
	// 		fmt.Println(result)
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"status":  "success",
	// 		"data": 1,
	// 	})
	// })

	// Grouping routes v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login",  func(c *gin.Context) {
			type Login struct {
				Email string `form:"email" json:"email" binding:"required"`
				Password string `form:"password" json:"password" binding:"required"`
			}
	
			var json2 Login
			if err := c.ShouldBindJSON(&json2); err == nil {
				c.JSON(200, gin.H{
					"messages": "inserted v1",
					"data": json2,
				})
		
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		})
	}

	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	r.GET("/user/:id/pt/:ptid", func(c *gin.Context) {
		id := c.Param("id")
		ptid := c.Param("ptid")
		fmt.Println("aaaa", id, ptid)
		c.String(http.StatusOK, "Hello %s", id)
	})

	// r.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(http.StatusOK, message)
	// })

	r.GET("/users/:name", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/form-post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		
		type CreatePost struct {
			Title string `form:"title" json:"title" binding:"required"`
			Body string `form:"body" json:"body" binding:"required"`
		}

		var json CreatePost

		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(200, gin.H{
				"messages": "inserted",
				"data": json,
			})
	
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	})

	// Map as querystring or postform parameters
	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	// Content-Type: application/x-www-form-urlencoded
    // names[first]=thinkerou&names[second]=tianou
	r.POST("/post-map", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	r.Run(":3000")
}
