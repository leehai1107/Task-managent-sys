package handler

import (
	"net/http"

	"scheduleSystem/model/entity"
	"scheduleSystem/model/request"
	"scheduleSystem/usecase"
	"scheduleSystem/wapper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IUserHandler interface {
	Ping(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHander struct {
	userService usecase.IUserService
}

func NewUserHandler() IUserHandler {
	return &userHander{
		userService: usecase.NewUserService(),
	}
}

func (u *userHander) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData("Pong..."))
}

func (u *userHander) Register(c *gin.Context) {
	var req entity.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
	}

	res, err := u.userService.Register(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
	} else {
		c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
	}
}

func (u *userHander) Login(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var req request.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
	}

	res, err := u.userService.Login(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
	} else {
		c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
	}
}
