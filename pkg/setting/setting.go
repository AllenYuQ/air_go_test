package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	mylog "github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize      int
	IdentityKey   string
	RocketAddress string

	LogFilePath = "./logs"
	LogFileName = "air"
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	initLogger()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	RocketAddress = Cfg.Section("rocketmq").Key("ROCKET_ADDRESS").MustString("127.0.0.1:9876")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	IdentityKey = sec.Key("IDENTITY_KEY").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func initLogger() {
	filePath := "./logs/air.log-%Y%m%d"
	fileWriter, err := rotatelogs.New(
		filePath,
		rotatelogs.WithLinkName("./logs/air.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*7),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)
	if err != nil {
		fmt.Println("初始化日志失败")
		os.Exit(1)
	}
	mylog.SetOutput(io.MultiWriter(os.Stdout, fileWriter))
	//设置日志的格式
	mylog.SetFormatter(&mylog.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   true,
		FullTimestamp:   true,
	})
	mylog.SetReportCaller(true)
	mylog.SetLevel(mylog.DebugLevel)
}
