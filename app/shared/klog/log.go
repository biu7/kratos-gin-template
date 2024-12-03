package klog

import (
	"github.com/biu7/gokit/log"
	"os"

	"github.com/spf13/cast"

	klog "github.com/go-kratos/kratos/v2/log"
)

// Logger Kratos log wrapper, only for Kratos.
type Logger struct {
	logger log.Logger
}

func NewLogger(logger log.Logger) klog.Logger {
	return &Logger{
		logger: logger,
	}
}

func SetKratosDefaultLogger(logger log.Logger) {
	klog.SetLogger(NewLogger(logger))
}

func (l *Logger) args(args []interface{}) (string, []interface{}) {
	var msg string
	var ret []interface{}
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			ret = append(ret, args[i])
			break
		}
		if args[i] == klog.DefaultMessageKey {
			msg = cast.ToString(args[i+1])
			continue
		}
		ret = append(ret, args[i], args[i+1])
	}

	return msg, ret
}

func (l *Logger) Log(level klog.Level, args ...interface{}) error {
	var msg string
	msg, args = l.args(args)
	switch level {
	case klog.LevelDebug:
		l.logger.Debug(msg, args...)
	case klog.LevelInfo:
		l.logger.Info(msg, args...)
	case klog.LevelWarn:
		l.logger.Warn(msg, args...)
	case klog.LevelError:
		l.logger.Error(msg, args...)
	case klog.LevelFatal:
		args = append([]interface{}{
			"level", "FATAL",
		}, args...)
		l.logger.Error(msg, args...)
		os.Exit(1) //nolint: revive
	default:
		args = append([]interface{}{
			"level", "FATAL",
		}, args...)
		l.logger.Info(msg, args...)
	}
	return nil
}
