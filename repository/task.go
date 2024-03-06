package repository

import (
	"context"

	"github.com/leehai1107/Task-managent-sys/constant"
	"github.com/leehai1107/Task-managent-sys/infra"
	"github.com/leehai1107/Task-managent-sys/model/request"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"go.uber.org/zap"
)

type ITaskRepo interface {
	CreateTask(ctx context.Context, req request.Task) (res response.TaskResponse, err error)
}

type taskRepo struct{}

func NewTaskRepo() ITaskRepo {
	return &taskRepo{}
}

func (t *taskRepo) CreateTask(ctx context.Context, req request.Task) (res response.TaskResponse, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Create(&req).Error
	if err != nil {
		return response.TaskResponse{Title: "", Message: constant.FAIL_MESSAGE}, err
	}
	return response.TaskResponse{Title: req.Title, Message: constant.SUCCESS_MESSAGE}, nil
}
