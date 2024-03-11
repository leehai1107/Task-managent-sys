package usecase

import (
	"context"

	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"github.com/leehai1107/Task-managent-sys/repository"
	"github.com/leehai1107/Task-managent-sys/utils"
	"github.com/leehai1107/Task-managent-sys/wapper"
	"go.uber.org/zap"
)

type ITaskService interface {
	CreateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error)
}

type taskService struct {
	taskRepo repository.ITaskRepo
}

func NewTaskService() ITaskService {
	return &taskService{
		taskRepo: repository.NewTaskRepo(),
	}
}

func (t *taskService) CreateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = utils.ValidateDate(req.StartDate, req.EndDate)
	if err != nil {
		return response.TaskResponse{}, wapper.Wrap(err, "ValidateDate")
	}

	res, err = t.taskRepo.CreateTask(ctx, req)
	if err != nil {
		return response.TaskResponse{}, err
	}

	return res, nil
}