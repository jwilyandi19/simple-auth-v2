package user

import (
	"context"

	"github.com/jwilyandi19/simple-auth-v2/domain/user"
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
	return nil, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*user.User, error) {
	return nil, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]user.User, error) {
	return nil, nil
}

func (u *userUsecase) Update(ctx context.Context, id string, user *user.User) (*user.User, error) {
	return nil, nil
}
