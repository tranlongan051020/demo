package user

import (
	"demo/constant"
	userEntity "demo/internal/entity/user"
	"demo/pkg/aerror"
	"demo/pkg/db"

	"context"
)

func (ist *instance) Update(ctx context.Context, userId uint32, data userEntity.UserUpdate) error {
	query := db.Get(ctx).Table(constant.UserTableName).Where("id", userId).Updates(data)
	if query.Error != nil {
		err := aerror.New(ctx, aerror.ErrUnexpected)
		return err
	}

	return nil
}
