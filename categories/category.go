package categories

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
	"strconv"
)

type responseModel struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	CreatedBy   int    `json:"created_by"`
}

func AddNewCategory(c *gin.Context) {
	var category models.Category
	category.Name = c.PostForm("name")
	category.Description = c.PostForm("desc")
	userId := c.PostForm("created_by")
	category.UserID, _ = strconv.Atoi(userId)
	database.DB.Create(&category)
	c.JSON(200, &responseModel{
		Name:        category.Name,
		Description: category.Description,
		CreatedBy:   category.UserID,
	})
}
