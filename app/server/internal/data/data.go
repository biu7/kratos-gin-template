package data

import (
	"github.com/biu7/gokit/log"
	"kratos-gin-template/app/server/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
