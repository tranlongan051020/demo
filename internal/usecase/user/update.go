package user

import (
	"context"
	"errors"

	userEntity "demo/internal/entity/user"
	"demo/pkg/aerror"
	"demo/pkg/validator"
)

type UpdateInput struct {
	UserId     uint32 `json:"user_id" validate:"required,gte=1,lte=9999"`
	UpdateData userEntity.UserUpdate
}

func (ist *instance) Update(ctx context.Context, input UpdateInput) error {
	errs := validator.GetValidator().ValidateStruct(ctx, input)
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	user, err := ist.userSto.Get(ctx, input.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		err = aerror.New(ctx, ErrorUserDoesNotExist)
		return err
	}

	err = ist.userSto.Update(ctx, input.UserId, input.UpdateData)
	if err != nil {
		return err
	}

	return nil
}
