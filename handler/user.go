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

// Ping godoc
// @Summary      Ping service
// @Description  get ping from service
// @Tags         ping
// @Success      200  {object}   wapper.Response
// @Router       /user/ping [get]
func (u *userHander) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData("Pong...Pong...Pong"))
}

// Register godoc
// @Summary      Register user
// @Description  Register for user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 data body entity.User true "user register data"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /user/register [post]
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

// Login godoc
// @Summary      Login user
// @Description  Login for user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 data body request.UserLogin true "user login data"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /user/login [post]
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

// UpdateUser godoc
// @Summary      Update user
// @Description  Update infomations for user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 data body entity.User true "user update data"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /user/update [put]
func (u *userHander) UpdateUser(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	var req entity.User
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
