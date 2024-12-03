package main

import (
	"github.com/biu7/gokit/env"
	"github.com/biu7/gokit/log"
	"kratos-gin-template/app/server/internal/conf"
	"kratos-gin-template/app/shared/klog"

	_ "go.uber.org/automaxprocs"
)

func main() {
	logger := initLogger()
	bc := initConfig()
	runApp(bc, logger)
}

func initLogger() log.Logger {
	logLevel := log.LevelInfo
	if env.Debug() || env.Local() {
		logLevel = log.LevelDebug
	}
	newLogger := log.With(
		log.NewLogger(logLevel),
		"caller", log.Caller(4),
		"version", Version,
		"traceID", log.TraceID(),
		"spanID", log.SpanID(),
	)
	klog.SetKratosDefaultLogger(newLogger)
	return newLogger
}

func initConfig() *conf.Bootstrap {
	bootstrap, err := conf.Load("")
	if err != nil {
		panic(err)
	}
	return bootstrap
}
