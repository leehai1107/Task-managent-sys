package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/Task-managent-sys/model/request"
	"github.com/leehai1107/Task-managent-sys/usecase"
	"github.com/leehai1107/Task-managent-sys/wapper"
)

type ITaskHandler interface {
	CreateTask(c *gin.Context)
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
	/* BUG: did not check the enddate and startdate: enddate > startdate   */
	var req request.Task
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
