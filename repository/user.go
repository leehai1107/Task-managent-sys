package repository

import (
	"context"

	"github.com/leehai1107/Task-managent-sys/infra"
	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"go.uber.org/zap"
)

type IUserRepo interface {
	Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error)
	GetUserByUsername(ctx context.Context, req string) (res entity.User, err error)
	Save(ctx context.Context, req entity.User) (res entity.User, err error)
}

type userRepo struct{}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (u *userRepo) Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Create(&req).Error
	if err != nil {
		// logger.Error("[RegisterRepo]err:", zap.Any("err", err))
		return response.UserCreateRes{UserID: 0}, err
	}
	// logger.Info("[Register]res:", zap.Any("res", res))
	return response.UserCreateRes{UserID: req.UserID}, nil
}

func (u *userRepo) GetUserByUsername(ctx context.Context, req string) (res entity.User, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Where("username = ?", req).First(&res).Error
	if err != nil {
		return entity.User{}, err
	}
	return
}

func (u *userRepo) Save(ctx context.Context, req entity.User) (res entity.User, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err = infra.GetDB().Updates(req).Scan(&res).Error
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}
