package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-online-store/users"
)

func UserRoute(router *gin.Engine) {
	routerGroup := router.Group("/user")
	routerGroup.Use(Middleware)
	routerGroup.GET("/", users.GetAllUsers)
	routerGroup.GET("/:user_id", users.GetUserById)
	routerGroup.POST("/", users.AddNewUser)
	routerGroup.PUT("/:user_id", users.UpdateUserData)
	routerGroup.DELETE("/:user_id", users.DeleteUserData)

}

func Middleware(c *gin.Context) {
	fmt.Println("Testing Middleware")
	c.JSON(200, gin.H{
		"data": "No message here",
	})
}
