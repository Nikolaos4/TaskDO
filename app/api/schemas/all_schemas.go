package schemas
import (
	"github.com/gin-gonic/gin"
	"time"
)

//структуры для пользователя
type CreateUserRequest struct {
	Login string `json:"login" binding:"required"`
	//Name string `json:"name" binding:"required"`
	//LastName string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUser struct {
	Login string `json:"login" binding:"required"`
	//Name string `json:"name" binding:"required"`
	//LastName string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
type ResponseUser struct {
	ID uint `json:"id"`
	Login string `json:"login"`
	//Name string `json:"name"`
	//LastName string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}
type ResponseToken struct {
	Token string `json:"token"`
}


//структуры для задач
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

type UpdateTask struct {
	Title string `json:"title"`
	Data time.Time `json:"data"`
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