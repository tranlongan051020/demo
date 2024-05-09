package user

import (
	"demo/constant"
	userEntity "demo/internal/entity/user"
	"demo/pkg/aerror"
	"demo/pkg/db"

	"context"
)

func (ist *instance) Get(ctx context.Context, userId uint32) (*userEntity.User, error) {
	data := new(userEntity.User)
	result := db.Get(ctx).
		Table(constant.UserTableName).
		Where("id", userId).
		Find(data)

	if result.Error != nil {
		err := aerror.New(ctx, aerror.ErrUnexpected)
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return data, nil
}
