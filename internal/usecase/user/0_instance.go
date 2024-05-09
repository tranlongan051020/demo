package user

import (
	"context"
	userSto "demo/internal/storage/user"
)

type UserUsecase interface {
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	Update(ctx context.Context, input UpdateInput) error
}

type instance struct {
	userSto userSto.UserStorage
}

func New(userSto userSto.UserStorage) UserUsecase {
	return &instance{
		userSto: userSto,
	}
}
