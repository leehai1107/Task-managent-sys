package usecase

import (
	"context"
	"errors"

	"github.com/leehai1107/Task-managent-sys/model/entity"
	"github.com/leehai1107/Task-managent-sys/model/request"
	"github.com/leehai1107/Task-managent-sys/model/response"
	"github.com/leehai1107/Task-managent-sys/repository"
	"github.com/leehai1107/Task-managent-sys/wapper"
	"go.uber.org/zap"
)

type IUserService interface {
	Register(ctx context.Context, req entity.User) (res response.UserCreateRes, err error)
	Login(ctx context.Context, req request.UserLogin) (res entity.User, err error)
	Update(ctx context.Context, req request.UserUpdateReq) (res entity.User, err error)
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

func (u *userService) Update(ctx context.Context, req request.UserUpdateReq) (res entity.User, err error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	res, err = u.userRepo.Save(ctx, req)
	if err != nil {
		return entity.User{}, wapper.Wrap(err, "Some thing wrong!")
	}

	return res, nil
}
