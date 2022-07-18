package userlogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-online-store/database"
	"go-online-store/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var user models.User

func LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := getUserByEmail(email)
	if err != nil {
		fmt.Println("error getting user by username, check username")
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": "Unauthorized, please check username or password",
		})
		c.Abort()
		return
	}
	passErr := getPassword(password)
	if passErr != nil {
		fmt.Println("password error")
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": "Unauthorized, please check username or password",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"data": "logged in successfully",
	})

}

func getUserByEmail(email string) error {
	err := database.DB.Debug().Model(&user).Where("email = ?", email).Scan(&user)
	fmt.Println(&user)
	if err.RowsAffected <= 0 {
		fmt.Println("getUserByUsername() failed, err:", err)
		return gorm.Errors{}
	}
	return nil
}

func getPassword(password string) error {
	hashedPass := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	fmt.Println("error from bcrypt:", err)
	if err != nil {
		return err
	}
	return nil
}
