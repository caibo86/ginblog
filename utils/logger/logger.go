package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:03:04",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// fileName := fmt.Sprintf("%s:%d", frame.File, frame.Line)
			fileName := fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
			return frame.Function, fileName
		},
	})
}

// GetLogger 日志句柄获取
func GetLogger() *logrus.Logger {
	return logger
}
