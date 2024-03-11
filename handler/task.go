package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/usecase"
	"github.com/leehai1107/Task-managent-sys/wapper"
)

type ITaskHandler interface{
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

/* TODO: Update task , update task base on date time  */
