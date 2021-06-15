package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 注册的简单方式 最基础的方式
func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	if err := r.Run(":9090");err != nil {
		fmt.Println("start network failed")
	}

}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "houzhenguo",
		"age": 18,
	})
}
