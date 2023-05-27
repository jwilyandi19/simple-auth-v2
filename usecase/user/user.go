package user

import (
	"context"

	"github.com/jwilyandi19/simple-auth-v2/domain/user"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	userRepository "github.com/jwilyandi19/simple-auth-v2/repository/user"
)

type UserUsecase interface {
	Create(ctx context.Context, user *user.User) (*user.User, error)
	GetByID(ctx context.Context, id string) (*user.User, error)
	GetAll(ctx context.Context) ([]user.User, error)
	Update(ctx context.Context, id string, user *user.User) (*user.User, error)
}

type userUsecase struct {
	userRepo userRepository.UserRepository
}

func NewUserUsecase(u userRepository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *user.User) (*user.User, error) {
	var err error
	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*user.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]user.User, error) {
	res, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *userUsecase) Update(ctx context.Context, id string, user *user.User) (*user.User, error) {
	var err error
	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	res, err := u.userRepo.Update(ctx, user, id)
	if err != nil {
		return res, err
	}
	return res, nil
}
