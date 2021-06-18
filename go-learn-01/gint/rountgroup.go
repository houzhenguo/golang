package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 路由组
func main() {
	r := gin.Default()
	r.Use(StatCost()) // 注册一个全局中间件
	userGroup := r.Group("/user")
	{
		userGroup.POST("/add", SetName(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "add one user",
			})
		})
		userGroup.GET("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "update one user",
			})
		})
		// 嵌套 group
		propGroup := userGroup.Group("/property")
		propGroup.GET("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "prop add one ",
			})
		})
		propGroup.GET("/update", func(c *gin.Context) {
			name := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{
				"message": "prop update one " + name,
			})
		})
		r.Run(":9090")
	}
}

// 定义中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("username", "侯振国") // 可以通过c.Set在请求上下文中设置值，在后续的处理函数能够取到这个值
		// 调用该请求的剩余处理程序
		c.Next()
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("cost time is ", cost)
	}
}
// 单独注册一个 中间件
func  SetName() gin.HandlerFunc {
	return func(c *gin.Context) {
		age:= c.PostForm("age")
		fmt.Println("age is", age)
	}
}