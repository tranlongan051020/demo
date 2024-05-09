package controller

import (
	userCtrl "demo/api/controller/user"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	New,

	userCtrl.New,
)
