package main
import (
	"github.com/gin-gonic/gin"
	"TaskDo/app/db"
	"TaskDo/app/api/handlers"
)

func main() {
	db.InitDB()
	r := gin.Default()
	//для работы с пользователем
	r.POST("/user", handlers.CreateUser)
	r.GET("/user/:id", handlers.GetUserID)
	//для работы с задачами
	r.POST("/task", handlers.CreateTask)
	r.GET("/task/:id", handlers.GetTask)
	//прослушка порта
	r.Run(":8080")
}