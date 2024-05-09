package storage

import (
	"demo/internal/storage/user"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	user.New,
)
