package handlers

import (
	"TaskDo/app/api/schemas"
	"TaskDo/app/db"
	"TaskDo/app/core"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req schemas.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		schemas.ResponseJSON(c, http.StatusBadRequest, "Invailed input", err)
		return
	}

	password, err := security.Password_hashing(req.Password)
	if err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "Faild to hash password", err)
		return
	}
	data := db.User{
		Login: req.Login,
		//Name: req.Name,
		//LastName: req.LastName,
		Password: password,
	}
	if err := db.DB.Create(&data).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "The user has not been created", err)
	}

	user_data := schemas.ResponseUser{
		ID: data.ID,
		Login: data.Login,
		//Name: data.Name,
		//LastName: data.LastName,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The user was successfully created:", user_data)
}

func LoginUser(c *gin.Context) {
	var req schemas.LoginUser
	if err := c.ShouldBind(&req); err != nil {
		schemas.ResponseJSON(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	current_user := db.User{Login: req.Login} //Name: req.Name, LastName: req.LastName
	if err := db.DB.First(&current_user).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusNotFound, "User is not found", err)
		return
	}

	token, err := security.GenerateToken(current_user.ID)
	if err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "Unable to create a token", err)
		return
	}
	data := schemas.ResponseToken{
		Token: token,
	}
	schemas.ResponseJSON(c, http.StatusOK, "Successful entry!", data)
}

func GetUserID(c *gin.Context) {
	var current_user db.User
	if err := db.DB.First(&current_user, c.Param("id")).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusNotFound, "User is not found", nil)
		return
	}
	user_data := schemas.ResponseUser{
		ID: current_user.ID,
		Login: current_user.Login,
		//Name: current_user.Name,
		//LastName: current_user.LastName,
		CreatedAt: current_user.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The user was successfully found", user_data)
}
