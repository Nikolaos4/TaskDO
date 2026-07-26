package schemas
import (
	"github.com/gin-gonic/gin"
	"time"
)

//структура ввода и вывода данных при создании пользователя 
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type ResponseUser struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	LastName string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

//структура для создания задачи
type CreateTask struct {
	Title string `json:"title" binding:"required"`
	Data time.Time `json:"data" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

type ResponseTask struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Data time.Time `json:"data"`
	UserID uint `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

//
type JsonResponse struct {
	Status int
	Message string
	Data any
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status: status,
		Message: message,
		Data: data,
	}
	c.JSON(status, response)
}