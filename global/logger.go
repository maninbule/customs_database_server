package global

import (
	"github.com/customs_database_server/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	Logger *logger.Logger
)

func InitLogger() {
	fileName := AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}
