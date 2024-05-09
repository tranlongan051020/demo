package usecase

import (
	"demo/internal/usecase/user"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	user.New,
)
