package main
import (
	"github.com/gin-gonic/gin"
	"TaskDo/app/db"
	"TaskDo/app/api/handlers"
	"TaskDo/app/midleware"
)

func main() {
	db.InitDB()
	r := gin.Default()

	//для работы с пользователем
	r.POST("/user/create", handlers.CreateUser)
	r.POST("/user/login", handlers.LoginUser)
	r.GET("/user/:id", handlers.GetUserID)

	//для работы с задачами
	r.GET("/task/:id", midleware.AuthenticateMiddleware, handlers.GetTask)
	r.GET("/task/all_tasks", midleware.AuthenticateMiddleware, handlers.GetUserTasks)
	r.POST("/task/create", midleware.AuthenticateMiddleware, handlers.CreateTask)
	r.PUT("/task/:id", midleware.AuthenticateMiddleware, handlers.UpdateTask)
	r.DELETE("/task/:id", midleware.AuthenticateMiddleware, handlers.DeleteTask)

	//прослушка порта
	r.Run(":8080")
}