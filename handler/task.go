package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/usecase"
	"github.com/leehai1107/Task-managent-sys/wapper"
)

type ITaskHandler interface {
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	GetTaskByUserID(c *gin.Context)
	GetExpiredTask(c *gin.Context)
}

type taskHandler struct {
	taskService usecase.ITaskService
}

func NewTaskHandler() ITaskHandler {
	return &taskHandler{
		taskService: usecase.NewTaskService(),
	}
}

func (t *taskHandler) CreateTask(c *gin.Context) {
	var req entity.Task
	err := c.BindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := t.taskService.CreateTask(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}

	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}

func (t *taskHandler) UpdateTask(c *gin.Context) {
	var req entity.Task
	err := c.BindJSON(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := t.taskService.UpdateTask(c, req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}

func (t *taskHandler) GetTaskByUserID(c *gin.Context) {
	req := c.Param("user_id")
	data, err := strconv.Atoi(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := t.taskService.GetTaskByUserID(c, uint(data))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}

func (t *taskHandler) GetExpiredTask(c *gin.Context) {
	req := c.Param("user_id")
	data, err := strconv.Atoi(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := t.taskService.GetExpiredTask(c, uint(data))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}
