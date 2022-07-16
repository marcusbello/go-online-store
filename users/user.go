package users

import (
	"github.com/gin-gonic/gin"
	"go-online-store/database"
	"go-online-store/models"
)

type responseModel struct {
	UserName  string `json:"user_name"`
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
			UserName:  get.UserName,
			Email:     get.Email,
			FirstName: get.Profile.FirstName,
			LastName:  get.Profile.LastName,
			UserRole:  get.UserRole.Name,
		}
		resp = append(resp, response)
	}

	c.JSON(200, &resp)
}

func AddNewUser(c *gin.Context) {
	var users models.User
	users.UserName = c.PostForm("user_name")
	users.Email = c.PostForm("email")
	users.Password = c.PostForm("password")
	database.DB.Create(&users)
	c.JSON(200, &responseModel{
		UserName:  users.UserName,
		Email:     users.Email,
		FirstName: users.Profile.FirstName,
		LastName:  users.Profile.LastName,
		UserRole:  users.UserRole.Name,
	})
}

func GetUserById(c *gin.Context) {

}

func UpdateUserData(c *gin.Context) {

}

func DeleteUserData(c *gin.Context) {

}
