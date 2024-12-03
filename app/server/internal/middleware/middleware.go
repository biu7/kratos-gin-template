package middleware

import (
	"github.com/biu7/gokit/ginutils"
	"github.com/biu7/gokit/log"
	"github.com/google/wire"
	"kratos-gin-template/app/server/internal/conf"
)

var ProviderSet = wire.NewSet(NewMiddleware)

type Middleware struct {
	Auth *Auth
	Gin  *ginutils.Middleware
}

func NewMiddleware(c *conf.Biz, logger log.Logger) *Middleware {
	return &Middleware{
		Auth: &Auth{
			jwtSecret: c.GetSecret().GetJwtSecret(),
			log:       logger,
		},
		Gin: ginutils.NewMiddleware(logger),
	}
}
