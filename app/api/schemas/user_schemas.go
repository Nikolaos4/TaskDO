package schemas
import (
	"github.com/gin-gonic/gin"
)

//структуры для пользователя
type CreateUserRequest struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUser struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type ResponseUserData struct {
	ID uint `json:"id"`
	Login string `json:"login"`
	CreatedAt string `json:"created_at"`
}

type ResponseToken struct {
	Token string `json:"token"`
}


//структура для общего ответа
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