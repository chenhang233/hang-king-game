//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/topfreegames/pitaya/v2"
	"hang-king-game/app/batt/internal/biz"
	"hang-king-game/app/batt/internal/conf"
	"hang-king-game/app/batt/internal/data"
	"hang-king-game/app/batt/internal/server"
	"hang-king-game/app/batt/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*pitaya.Pitaya, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp2))
}
