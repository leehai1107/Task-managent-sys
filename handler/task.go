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
	GetAvailableTask(c *gin.Context)
}

type taskHandler struct {
	taskService usecase.ITaskService
}

func NewTaskHandler() ITaskHandler {
	return &taskHandler{
		taskService: usecase.NewTaskService(),
	}
}

// CreateTask godoc
// @Summary      Create task
// @Description  Create task which user_id = ?
// @Tags         task
// @Accept       json
// @Produce      json
// @Param 		 data body entity.Task true "task create data"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /task/create [post]
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

// UpdateTask godoc
// @Summary      Update task
// @Description  Update task data
// @Tags         task
// @Accept       json
// @Produce      json
// @Param 		 data body entity.Task true "task update data"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /task/update [put]
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

// GetTaskByUserID godoc
// @Summary      Get tasks by user_id
// @Description  Get all tasks by user_id
// @Tags         task
// @Param 		 user_id path int true "User ID"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /task/{user_id} [get]
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

// GetExpiredTask godoc
// @Summary      Get tasks expired
// @Description  Get all tasks end_time < now()
// @Tags         task
// @Param 		 user_id path int true "User ID"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /task/expired/{user_id} [get]
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

// GetAvailableTask godoc
// @Summary      Get tasks available
// @Description  Get all tasks end_time > now()
// @Tags         task
// @Param 		 user_id path int true "User ID"
// @Success      200  {object}   wapper.Response
// @Failure	   400 {object} wapper.Response
// @Failure		500 {object} wapper.Response
// @Router       /task/available/{user_id} [get]
func (t *taskHandler) GetAvailableTask(c *gin.Context) {
	req := c.Param("user_id")
	data, err := strconv.Atoi(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, wapper.FailWithErr(err))
		return
	}
	res, err := t.taskService.GetAvailableTask(c, uint(data))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wapper.FailWithErr(err))
		return
	}
	c.IndentedJSON(http.StatusOK, wapper.SuccessWithData(res))
}
