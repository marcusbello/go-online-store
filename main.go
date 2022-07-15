package main

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/routes"
)

func main() {

	router := gin.New()
	database.DBConnect()
	routes.UserRoute(router)
	routes.ProductRoute(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello World",
		})
	})
	router.Run(":8000")
}
