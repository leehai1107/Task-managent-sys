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
}

func NewRouter() IRouter {
	return &router{
		userHandler: handler.NewUserHandler(),
	}
}

func (r *router) Register() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/ping", r.userHandler.Ping)
		user.POST("/register", r.userHandler.Register)
		user.POST("/login", r.userHandler.Login)
	}

	return router
}
