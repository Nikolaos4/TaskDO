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
		schemas.ResponseJSON(c, http.StatusBadRequest, "Invailed input", nil)
		return
	}

	password, err := security.Password_hashing(req.Password)
	if err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "Faild to hash password", err)
		return
	}
	data := db.User{
		Name: req.Name,
		LastName: req.LastName,
		Password: password,
	}
	if err := db.DB.Create(&data).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "The user has not been created", err)
	}

	user_data := schemas.ResponseUser{
		ID: data.ID,
		Name: data.Name,
		LastName: data.LastName,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The user was successfully created:", user_data)
}


func GetUserID(c *gin.Context) {
	var current_user db.User
	if err := db.DB.First(&current_user, c.Param("id")).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusNotFound, "User is not found", nil)
		return
	}
	user_data := schemas.ResponseUser{
		ID: current_user.ID,
		Name: current_user.Name,
		LastName: current_user.LastName,
		CreatedAt: current_user.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The user was successfully found", user_data)
}
