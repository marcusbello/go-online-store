package routes

import (
	"github.com/gin-gonic/gin"
	"go-online-store/products"
)

func ProductRoute(router *gin.Engine) {
	router.GET("/product", products.GetAllProducts)
	router.POST("/product", products.AddProduct)
}
