package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(costHandler)
	r.GET("/index", handleIndex)
	r.GET("/json", handleJson)
	loadShopGroup(r)
	r.Run()
}

// 嵌套 book 并且 加了针对book的过滤条件
func loadShopGroup(r *gin.Engine) {
	bookGroup := r.Group("/book")
	bookGroup.Use(bookFilter)
	{
		bookGroup.GET("/buy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "buy",
			})
		})
		bookGroup.GET("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "add",
			})
		})
	}
}

// 只是对于book 这个group 的分组
func bookFilter(c *gin.Context) {
	fmt.Println("book in")
	c.Next()
	fmt.Println("book out")
}
func handleIndex(c *gin.Context) {
	fmt.Println("in handle index")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
	fmt.Println("out handle index")
}

func handleJson(c *gin.Context) {
	fmt.Println("in handle json")
	name, ok := c.Get("name")
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello json",
			"name":    name,
		})
	}

	fmt.Println("out handle json")
}
func costHandler(c *gin.Context) {
	fmt.Println("in cost ")
	start := time.Now()
	c.Set("name", "houzhenguo")
	c.Next()
	cost := time.Since(start)
	c.JSON(200, gin.H{
		"cost": cost,
	})
	fmt.Println("cost end cost is", cost)
}
