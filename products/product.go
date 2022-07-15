package products

import "github.com/gin-gonic/gin"

func GetAllProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello World From Products",
	})
}
