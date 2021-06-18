package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetUpRouter 配置路由信息

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/hello", helloHandler)
	LoadBook(r)
	LoadShop(r)
	return r
}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":"zhangsan",
		"age":8,
	})
}
