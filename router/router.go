package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leehai1107/Task-managent-sys/handler"
)

type IRouter interface {
	Register() *gin.Engine
}

type router struct {
	userHandler handler.IUserHandler
	taskHandler handler.ITaskHandler
}

func NewRouter() IRouter {
	return &router{
		userHandler: handler.NewUserHandler(),
		taskHandler: handler.NewTaskHandler(),
	}
}

func (r *router) Register() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/ping", r.userHandler.Ping)
		user.POST("/register", r.userHandler.Register)
		user.POST("/login", r.userHandler.Login)
		user.PUT("/update", r.userHandler.UpdateUser)
	}

	task := router.Group("/task")
	{
		task.POST("/create", r.taskHandler.CreateTask)
		task.PUT("/update", r.taskHandler.UpdateTask)
		task.GET("/:user_id", r.taskHandler.GetTaskByUserID)
	}

	return router
}
