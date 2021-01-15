package test

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var logger *logrus.Entry
var logger1 *logrus.Entry

func init() {
	//设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logger = logrus.WithFields(logrus.Fields{"request_id": "123444"})
	logger1 = logrus.WithFields(logrus.Fields{"user_ip": "127.0.0.1"})
}

func TestLog(t *testing.T) {
	logger.Info("hello, logrus....")
	logger1.Debug("hello, logrus1....")
}
