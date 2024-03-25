package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

type ComLogger struct {
	Log *logrus.Logger
}

var Logger = New()

func New() *ComLogger {
	var ginLog = logrus.New()
	ginLog.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.
	ginLog.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
	}
	return &ComLogger{ginLog}
}

func (logger *ComLogger) Info(args ...interface{}) {
	logger.Log.Info(args)
}

func (logger *ComLogger) Error(args ...interface{}) {
	logger.Log.Error(args)
}
