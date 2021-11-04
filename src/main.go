package main

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HelloGin",
	})
}

func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Posted",
	})
}

func QueryParameters(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	server := gin.Default()

	server.GET("/", Home)
	server.POST("/post", Post)
	server.GET("/query", QueryParameters)
	server.GET("/path/:name/:age", PathParameters)

	server.Run()
}
