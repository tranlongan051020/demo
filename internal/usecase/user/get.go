package user

import (
	"context"
	"errors"

	userEntity "demo/internal/entity/user"
	"demo/pkg/aerror"
	"demo/pkg/validator"
)

type GetInput struct {
	UserId uint32 `json:"user_id" validate:"required,gte=1,lte=9999"`
}

type GetOutput struct {
	UserData *userEntity.User
	// Additional user data
}

func (ist *instance) Get(ctx context.Context, input GetInput) (*GetOutput, error) {
	errs := validator.GetValidator().ValidateStruct(ctx, input)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	user, err := ist.userSto.Get(ctx, input.UserId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		err = aerror.New(ctx, ErrorUserDoesNotExist)
		return nil, err
	}

	output := &GetOutput{
		UserData: user,
	}
	return output, nil
}
