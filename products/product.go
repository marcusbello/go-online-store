package products

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(200, &products)
}

func AddProduct(c *gin.Context) {
	var product models.Product
	product.Name = c.PostForm("name")
	product.CreatedBy = c.PostForm("created_by")
	database.DB.Create(&product)
	c.JSON(200, &product)
}
