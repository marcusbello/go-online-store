package routes

import (
	"github.com/gin-gonic/gin"
	"go-online-store/categories"
)

func CategoryRoute(router *gin.Engine) {
	router.POST("/category", categories.AddNewCategory)
}
