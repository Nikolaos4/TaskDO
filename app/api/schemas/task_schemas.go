package schemas
import (
	"time"
)

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