package db

import (
	"context"

	"gorm.io/gorm"
)

var instance *gorm.DB

func Get(ctx context.Context) *gorm.DB {
	return instance
}
