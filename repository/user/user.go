package user

import (
	"context"

	domain "github.com/jwilyandi19/simple-auth-v2/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User, id string) (*domain.User, error)
}

type userRepository struct {
	db mongo.Database
}

func NewUserRepository(db mongo.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User, id string) (*domain.User, error) {
	return nil, nil
}
