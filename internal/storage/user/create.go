package user

import (
	"demo/constant"
	userEntity "demo/internal/entity/user"
	"demo/pkg/aerror"
	"demo/pkg/db"

	"context"
)

func (ist *instance) Create(ctx context.Context, data []*userEntity.UserCreate) error {
	query := db.Get(ctx).
		Table(constant.UserTableName).
		CreateInBatches(data, constant.DefaultBatchSizeCreate)
	if query.Error != nil {
		err := aerror.New(ctx, aerror.ErrUnexpected)
		return err
	}

	return nil
}
