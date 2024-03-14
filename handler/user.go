package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/model/request"
	"github.com/leehai1107/Task-managent-sys/usecase"
	"github.com/leehai1107/Task-managent-sys/wapper"
	"go.uber.org/zap"
)

type IUserHandler interface {
	Ping(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	UpdateUser(c *gin.Context)
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
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData("Pong...Pong...Pong"))
}

func (u *userHander) Register(c *gin.Context) {
	var req entity.User
	err := c.BindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}

	res, err := u.userService.Register(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}

func (u *userHander) Login(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var req request.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}

	res, err := u.userService.Login(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}

func (u *userHander) UpdateUser(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	var req request.UserUpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := u.userService.Update(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))

}
