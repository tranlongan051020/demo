//go:build wireinject
// +build wireinject

package main

import (
	httpserver "demo/cmd/http_server"

	"github.com/google/wire"
)

func initServer() *httpserver.Server {
	wire.Build(WireSet)
	return &httpserver.Server{}
}
