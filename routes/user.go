package routes

import (
	"github.com/gin-gonic/gin"
	"go-online-store/users"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user", users.GetAllUsers)
}
