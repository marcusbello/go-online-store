package products

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
)

type responseModel struct {
	Name        string            `json:"name"`
	Description string            `json:"desc"`
	Price       string            `json:"price"`
	Category    []models.Category `json:"categories"`
	CreatedBy   int               `json:"created_by"`
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(200, &products)
}

func AddNewProduct(c *gin.Context) {
	var product models.Product
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("desc")
	product.Price = c.PostForm("price")
	database.DB.Create(&product)

	c.JSON(200, &responseModel{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    nil,
		CreatedBy:   0,
	})
}
