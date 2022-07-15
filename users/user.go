package users

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello World From User page",
	})
}
