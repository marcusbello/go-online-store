package users

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type responseModel struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserRole  string `json:"user_role"`
}

func GetAllUsers(c *gin.Context) {
	var resp []responseModel
	var users []models.User
	database.DB.Find(&users)
	data := users
	for _, get := range data {
		response := responseModel{
			Email:     get.Email,
			FirstName: get.Profile.FirstName,
			LastName:  get.Profile.LastName,
			UserRole:  get.UserRole,
		}
		resp = append(resp, response)
	}

	c.JSON(200, &resp)
}

func AddNewUser(c *gin.Context) {
	var user models.User
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	hashed, _ := HashPassword(password)
	user.Password = hashed
	database.DB.Create(&user)
	c.JSON(200, &responseModel{
		Email:     user.Email,
		FirstName: user.Profile.FirstName,
		LastName:  user.Profile.LastName,
		UserRole:  user.UserRole,
	})
}

func GetUserById(c *gin.Context) {
	var user models.User
	userId := c.Param("user_id")
	database.DB.First(&user, userId)
	c.JSON(200, &responseModel{
		Email:     user.Email,
		FirstName: user.Profile.FirstName,
		LastName:  user.Profile.LastName,
		UserRole:  user.UserRole,
	})
}

func UpdateUserData(c *gin.Context) {
	userId := c.Param("user_id")
	var user models.User
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	hashed, _ := HashPassword(password)
	user.Password = hashed
	database.DB.Debug().Model(&user).Where("ID = ?", userId).Updates(&user)
	c.JSON(200, &responseModel{
		Email:     user.Email,
		FirstName: user.Profile.FirstName,
		LastName:  user.Profile.LastName,
		UserRole:  user.UserRole,
	})
}

func DeleteUserData(c *gin.Context) {
	var user models.User
	userId := c.Param("user_id")
	database.DB.First(&user, "id = ?", userId).Delete(&user)
	database.DB.Delete(&user, userId)
	c.JSON(http.StatusNoContent, gin.H{
		"data": "No content",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(bytes), err
}
