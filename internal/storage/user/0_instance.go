package user

import (
	"context"

	userEntity "demo/internal/entity/user"
)

type UserStorage interface {
	Get(ctx context.Context, userId uint32) (*userEntity.User, error)
	Create(ctx context.Context, data []*userEntity.UserCreate) error
	Update(ctx context.Context, userId uint32, data userEntity.UserUpdate) error
}

type instance struct{}

func New() UserStorage {
	return &instance{}
}
