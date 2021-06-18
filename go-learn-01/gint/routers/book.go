package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBook(r *gin.Engine) {
	r.GET("/book/read", readBook)
	r.GET("/book/buy", buyBook)
}
// http://localhost:9090/book/read
func readBook(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message":"read a book",
	})
}
// http://localhost:9090/book/buy
func buyBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"buy book",
	})
}
