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
	r.GET("/task/:id", handlers.GetTask)
	r.POST("/task", handlers.CreateTask)
	r.PUT("/task/:id", handlers.UpdateTask)
	r.DELETE("/task/:id", handlers.DeleteTask)

	//прослушка порта
	r.Run(":8080")
}