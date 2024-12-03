package main

import (
	"github.com/biu7/gokit/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-gin-template/app/server/internal/conf"
	"kratos-gin-template/app/server/internal/constants"
	"kratos-gin-template/app/shared/klog"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
)

var (
	hostname, _ = os.Hostname()
)

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(hostname),
		kratos.Name(constants.APPNameEn),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(klog.NewLogger(logger)),
		kratos.Server(
			hs,
		),
	)
}

func runApp(bc *conf.Bootstrap, logger log.Logger) {
	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Biz, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
