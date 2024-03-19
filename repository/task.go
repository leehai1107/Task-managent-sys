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
	UpdateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error)
	GetByUserID(ctx context.Context, req uint) (res []entity.Task, err error)
	GetExpiredTask(ctx context.Context, req uint) (res []entity.Task, err error)
	GetAvailableTask(ctx context.Context, req uint) (res []entity.Task, err error)
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

func (t *taskRepo) UpdateTask(ctx context.Context, req entity.Task) (res response.TaskResponse, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Updates(&req).Error
	if err != nil {
		return response.TaskResponse{TaskID: 0, Title: "", Message: constant.FAIL_UPDATE_MESSAGE}, err
	}

	return response.TaskResponse{TaskID: req.TaskID, Title: req.Title, Message: constant.SUCCESS_UPDATE_MESSAGE}, nil
}

func (t *taskRepo) GetByUserID(ctx context.Context, req uint) (res []entity.Task, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().
		Table(entity.Task{}.TableName()).
		Select("*").
		Where("user_id = ?", req).
		Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *taskRepo) GetExpiredTask(ctx context.Context, req uint) (res []entity.Task, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().
		Table(entity.Task{}.TableName()).
		Where("end_date < now() AND user_id = ?", req).
		Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *taskRepo) GetAvailableTask(ctx context.Context, req uint) (res []entity.Task, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().
		Table(entity.Task{}.TableName()).
		Where("end_date > now() AND user_id = ?", req).
		Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
