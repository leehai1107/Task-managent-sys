package usecase

import (
	"context"
	"errors"

	"scheduleSystem/model/entity"
	"scheduleSystem/model/request"
	"scheduleSystem/model/response"
	"scheduleSystem/repository"
	"scheduleSystem/wapper"

	"go.uber.org/zap"
)

type IUserService interface {
	Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error)
	Login(ctx context.Context, req request.UserLogin) (res entity.User, err error)
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
		// logger.Error("[Register]error:",zap.Any("err", err))
		return res, err
	}

	// logger.Info("[Register]res:",zap.Any("res", res))
	return res, nil
}

func (u *userService) Login(ctx context.Context, req request.UserLogin) (res entity.User, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	res, err = u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		// logger.Error("[GetUserByUsername]err", zap.Any("err", err))
		return entity.User{}, wapper.Wrap(err, "Wrong Username")
	}
	if req.Password != res.Password {
		return entity.User{}, wapper.Wrap(errors.New("Incorrect Username or Password"), "Wrong Password")
	}
	return res, nil
}
