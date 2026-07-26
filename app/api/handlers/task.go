package handlers

import (
	"TaskDo/app/api/schemas"
	"TaskDo/app/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var req schemas.CreateTask
	if err := c.ShouldBind(&req); err != nil {
		schemas.ResponseJSON(c, http.StatusBadRequest, "Invalid input", err)
		return
	}
	data := db.Task{
		Title:  req.Title,
		Data:   req.Data,
		UserID: req.UserID,
	}

	if err := db.DB.Create(&data).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusInternalServerError, "The task has not been created", err)
		return
	}

	task_data := schemas.ResponseTask{
		ID:        data.ID,
		Title:     data.Title,
		Data:      data.Data,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The task was successfully created", task_data)
}

func GetTask(c *gin.Context) {
	var current_task db.Task
	if err := db.DB.First(&current_task, c.Param("id")).Error; err != nil {
		schemas.ResponseJSON(c, http.StatusNotFound, "Task is not found", err)
		return
	}
	current_data := schemas.ResponseTask{
		ID:        current_task.ID,
		Title:     current_task.Title,
		Data:      current_task.Data,
		UserID:    current_task.UserID,
		CreatedAt: current_task.CreatedAt.Format(time.RFC3339),
	}
	schemas.ResponseJSON(c, http.StatusOK, "The task was successfully found", current_data)
}
