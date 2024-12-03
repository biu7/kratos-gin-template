package server

import (
	"github.com/biu7/gokit/env"
	"github.com/gin-gonic/gin"
	"kratos-gin-template/app/server/internal/biz"
	"kratos-gin-template/app/server/internal/conf"
	"kratos-gin-template/app/server/internal/middleware"
	"kratos-gin-template/app/server/internal/server/router"
	"kratos-gin-template/app/server/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewGinServer(c *conf.Server,
	mid *middleware.Middleware,
	greater *service.GreeterService,

	_ *biz.GreeterUsecase,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	if env.Prod() {
		gin.SetMode(gin.ReleaseMode)
	}
	ginRouter := gin.New()
	ginRouter.Use(
		mid.Gin.Log(),
		mid.Gin.Recovery(),
		mid.Gin.Cors(),
	)

	router.RegisterHelloRouter(ginRouter, mid, greater)

	srv.HandlePrefix("/", ginRouter)
	return srv
}
