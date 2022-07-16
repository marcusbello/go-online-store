package main

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
	"go-online-store/routes"
)

func AddCategory(c *gin.Context) {
	var category models.Category
	category.Name = c.PostForm("name")
	category.CreatedBy = c.PostForm("created_by")
	database.DB.Create(&category)
	c.JSON(200, &category)
}

func main() {

	router := gin.Default()
	database.DBConnect()
	routes.UserRoute(router)
	routes.ProductRoute(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello World",
		})
	})
	router.POST("/add_category", AddCategory)
	router.Run(":8000")
}
