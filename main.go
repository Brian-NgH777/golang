package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	goroutines "learn/custom-package/goroutines"
	models "learn/model"
	repoImpl "learn/repository/repository-impl"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ToBe       bool = false
	NumberInit uint = 111
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func init() {
	println("Main package initialized", NumberInit)
	connectDB()
}

func main() {
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
		v1.POST("/login", func(c *gin.Context) {
			type Login struct {
				Email    string `form:"email" json:"email" binding:"required"`
				Password string `form:"password" json:"password" binding:"required"`
			}

			var json2 Login
			if err := c.ShouldBindJSON(&json2); err == nil {
				c.JSON(200, gin.H{
					"messages": "inserted v1",
					"data":     json2,
				})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		})

		v1.GET("/goroutines", func(c *gin.Context) {
			go a()
			go cccc()
			fmt.Println("acccc")
			go b()
			goroutines.Channel()
			c.JSON(200, gin.H{
				"messages": "goroutines v1",
			})
		})

		v1.GET("/select-channels", func(c *gin.Context) {
			queue := make(chan int)
			done := make(chan bool)

			go func() {
				for i := 0; i < 10; i++ {
					queue <- i
				}
				done <- true
			}()

			for {
				select {
				case v := <-queue:
					fmt.Println(v)
				case <-done:
					fmt.Println("Done")
					c.JSON(200, gin.H{
						"messages": "Select Channels v1",
					})
				}
			}
		})

		v1.GET("/buffered-channels", func(c *gin.Context) {
			// buffered Channel
			ch := make(chan int, 2) // cho phép đẩy 2 giá trị vào channel ch
			ch <- 100               // 1
			ch <- 200               // 2
			// ch <- 300 // 3

			// close(ch) // when close channel thì giá trị khi <-ch là default của type (chan int) is 0
			// khi đã close channel thì ko thể gán giá trị cho nó
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			// fmt.Println(<-ch)
			// fmt.Println(<-ch)
			// fmt.Println(<-ch)
			// fmt.Println(<-ch)

			// buffered Channel info
			// + Nó ko có block như unbuffered channel
			// + Add 3 vào channel vào ch thì ko dc
			// + <-ch thêm thì ko có giá trị

			c.JSON(200, gin.H{
				"messages": "buffered Channels v1",
			})
		})

		v1.GET("/unbuffered-channels", func(c *gin.Context) {
			// UnBuffered Channel
			ch := make(chan int) // or ch := make(chan int, 0)
			go func() {
				ch <- 100 // 1 gán 100 vào channecl ch
				fmt.Println("Next")
			}()
			fmt.Println(<-ch) // 2 or result:= <- ch chờ giá trị channel ch
			fmt.Println("Done")
			// Unbuffered Channel info
			//	+ Nó sẽ bị block khi không có 1 or 2 không run tiếp
			//	+ nếu ko có 2 thì fmt.Println("Next") sẽ không dc run nó block tại 1 ko có ai biến nào nhận giá trị ch
			//	+ nếu ko có 1 mà có <-ch thì thì error deadlock thì ko có giá trị dc gán vào ch nên nó block

			c.JSON(200, gin.H{
				"messages": "UnBuffered Channels v1",
			})
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
			Body  string `form:"body" json:"body" binding:"required"`
		}

		var json CreatePost

		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(200, gin.H{
				"messages": "inserted",
				"data":     json,
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

	r.GET("/testing", func(c *gin.Context) {
		// tests := [6]int{2, 3, 5, 7, 11, 13}
		// for i, value := range tests {
		// 	fmt.Println(i, value)
		// }
		var aa = goroutines.ArraySliceType()
		fmt.Println(aa)

	})

	r.GET("/find-body-category", func(c *gin.Context) {
		result := repoImpl.NewBodycategoryRepo(Mongo.Client.
			Database("pttrainer")).
			FindBodycategories()

		c.JSON(200, gin.H{
			"messages": "inserted",
			"data":     result,
		})
	})

	r.POST("/body-category", func(c *gin.Context) {
		type body struct {
			Name        string `form:"name" json:"name" binding:"required"`
			Thumbnail   string `form:"thumbnail" json:"thumbnail"`
			Description string `form:"description" json:"description"`
		}

		var json body
		if err := c.ShouldBindJSON(&json); err == nil {
			bodycategory := models.Bodycategories{
				BcName:        json.Name,
				BcThumbnail:   json.Thumbnail,
				BcDescription: json.Description,
			}
			result := repoImpl.NewBodycategoryRepo(Mongo.Client.
				Database("pttrainer")).
				InsertBodycategory(bodycategory)

			c.JSON(200, gin.H{
				"messages": "body category v1",
				"data":     result,
			})
		} else {
			c.JSON(500, gin.H{
				"messages": "false",
			})
		}
	})

	r.Run(":3000")
}

func a() {
	fmt.Println("aaaaaaaaa")
}

func b() {
	n := 1
	for n < 5 {
		fmt.Println("n", n)
		n += 1
	}
}

func cccc() {
	n := 10
	for n < 15 {
		fmt.Println("n", n)
		n += 1
	}
}

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

func connectDB() *MongoDB {
	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://pttrainer:NmSJWnnBmApV5sEu@maincluster.gkfe6.mongodb.net/pttrainer?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect mongo db success")
	Mongo.Client = client
	return Mongo
}
