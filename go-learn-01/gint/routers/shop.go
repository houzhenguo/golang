package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadShop http://localhost:9090/shop/buy
func LoadShop(e *gin.Engine) {
	e.GET("/shop/buy", buyHandler)
}
//
func buyHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"by from shop",
	})
}
