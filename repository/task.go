package repository

import (
	"context"

	"github.com/leehai1107/Task-managent-sys/constant"
	"github.com/leehai1107/Task-managent-sys/infra"
	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"go.uber.org/zap"
)

type ITaskRepo interface {
	CreateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error)
}

type taskRepo struct{}

func NewTaskRepo() ITaskRepo {
	return &taskRepo{}
}

func (t *taskRepo) CreateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Create(&req).Error
	if err != nil {
		return response.TaskResponse{TaskID: 0, Title: "", Message: constant.FAIL_MESSAGE}, err
	}
	return response.TaskResponse{TaskID: req.TaskID, Title: req.Title, Message: constant.SUCCESS_MESSAGE}, nil
}
