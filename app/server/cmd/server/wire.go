//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/biu7/gokit/log"
	"kratos-gin-template/app/server/internal/biz"
	"kratos-gin-template/app/server/internal/conf"
	"kratos-gin-template/app/server/internal/data"
	"kratos-gin-template/app/server/internal/middleware"
	"kratos-gin-template/app/server/internal/server"
	"kratos-gin-template/app/server/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Biz, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(middleware.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
