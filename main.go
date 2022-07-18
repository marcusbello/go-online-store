package main

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/routes"
	"go-online-store/userlogin"
)

func main() {

	router := gin.Default()
	database.DBConnect()
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.CategoryRoute(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Online Store API Project",
		})
	})
	router.POST("/login", userlogin.LoginHandler)
	router.Run(":8000")
}
