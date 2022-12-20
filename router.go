package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type gin_router struct {
	router *gin.Engine
}

func timeThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
		fmt.Println("finished")
	}
}

var db = make(map[string]string)

func convertIntToString(number int) string {
	return strconv.Itoa(number)
}

func setupRouter() *gin_router {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("Executing goroutine")
		arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		ch := make(chan int, len(arr))
		go timeThree(arr, ch)

		for i := range ch {
			fmt.Println("RESULT", i)
		}

		fmt.Println("DONE")
		time.Sleep(time.Second)
		fmt.Println("5 SECONDS")

		c.String(http.StatusOK, "heyah")
	})

	r.GET("/hello/interpreter", func(c *gin.Context) {
		hello := helloWorld()

		c.String(http.StatusOK, hello)
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	ginRouter := gin_router{
		router: r,
	}

	return &ginRouter
}
