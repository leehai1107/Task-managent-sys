package usecase

import (
	"context"

	"github.com/leehai1107/Task-managent-sys/model/request"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"github.com/leehai1107/Task-managent-sys/repository"
	"go.uber.org/zap"
)

type ITaskService interface {
	CreateTask(ctx context.Context, req request.Task) (res response.TaskResponse, err error)
}

type taskService struct {
	taskRepo repository.ITaskRepo
}

func NewTaskService() ITaskService {
	return &taskService{
		taskRepo: repository.NewTaskRepo(),
	}
}

func (t *taskService) CreateTask(ctx context.Context, req request.Task) (res response.TaskResponse, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	res, err = t.taskRepo.CreateTask(ctx, req)
	if err != nil {
		return response.TaskResponse{}, err
	}

	return res, nil
}
