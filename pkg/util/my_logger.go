package util

import mylog "github.com/sirupsen/logrus"

var (
	DaoLog    = mylog.WithFields(mylog.Fields{"method": "dao"})
	PageLog   = mylog.WithFields(mylog.Fields{"method": "get/post"})
	RocketLog = mylog.WithFields(mylog.Fields{"method": "rocket"})
)
