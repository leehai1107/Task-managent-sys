package repository

import (
	"context"
	"errors"

	"scheduleSystem/infra"
	"scheduleSystem/model/entity"
	"scheduleSystem/model/response"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error)
	GetUserByUsername(ctx context.Context, req string) (res entity.User, err error)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// logger.Error("[ErrRecordNotFound]err", zap.Error(err))
			return entity.User{}, err
		}

		return entity.User{}, err
	}
	return
}
