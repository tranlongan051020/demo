package main

import (
	"demo/api/controller"
	"demo/api/router"
	httpserver "demo/cmd/http_server"
	"demo/internal/storage"
	"demo/internal/usecase"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	storage.WireSet,
	usecase.WireSet,
	controller.WireSet,
	router.New,
	httpserver.New,
)
