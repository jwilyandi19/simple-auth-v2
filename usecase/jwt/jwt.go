package jwt

import (
	"context"

	domain "github.com/jwilyandi19/simple-auth-v2/domain/user"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	"github.com/jwilyandi19/simple-auth-v2/repository/user"
	"github.com/labstack/echo/v4"
)

type JwtUsecase interface {
	SetJwtGeneral(g *echo.Group)
}

type jwtUsecase struct {
	UserRepo user.UserRepository
	Config   helper.Config
}

func NewJWTUsecase(u user.UserRepository, config helper.Config) JwtUsecase {
	return &jwtUsecase{
		UserRepo: u,
		Config:   config,
	}
}

func (j *jwtUsecase) getUserByID(ctx context.Context, id string) (*domain.User, error) {
	res, err := j.UserRepo.GetByID(ctx, id)
	if err != nil {
		return res, err
	}
	return res, nil
}
