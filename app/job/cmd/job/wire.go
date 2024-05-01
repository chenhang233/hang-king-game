//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"hang-king-game/app/job/internal/biz"
	"hang-king-game/app/job/internal/conf"
	"hang-king-game/app/job/internal/data"
	"hang-king-game/app/job/internal/server"
	"hang-king-game/app/job/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
