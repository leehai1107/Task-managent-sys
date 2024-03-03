package usecase

import (
	"context"

	"scheduleSystem/model/entity"
	"scheduleSystem/model/response"
	"scheduleSystem/repository"

	"go.uber.org/zap"
)

type IUserService interface {
	Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error)
}

type userService struct {
	userRepo repository.IUserRepo
}

func NewUserService() IUserService {
	return &userService{
		userRepo: repository.NewUserRepo(),
	}
}

func (u *userService) Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	res, err = u.userRepo.Register(ctx, req)
	if err != nil {
    logger.Error("[Register]error:",zap.Any("err", err))
		return res, err
	}

  logger.Info("[Register]res:",zap.Any("res", res))
	return res, nil
}
