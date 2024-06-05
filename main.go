package main

// @title 海关项目人脸识别、步态识别、高抗伪接口文档
// @version 1.0
// @description 第一个版本

// @host 172.21.116.147:8082
// @BasePath /

import (
	"fmt"
	"github.com/customs_database_server/global"
	Controllerkafka "github.com/customs_database_server/internal/controller/kafka"
	"github.com/customs_database_server/internal/model"
	"github.com/customs_database_server/internal/router"
)

func main() {
	global.SetupSetting() // 读取并解析配置文件
	global.InitDBEngine() // 初始化数据库
	global.InitLogger()   // 初始化日志记录器
	global.InitRedis()    // 初始化redis缓存
	fmt.Println("con nect database successful...version 4")
	model.InitModel() // 初始化数据库表
	go func() {
		Controllerkafka.GetImgFromKafka() // 不断从kafka接收图片，进行识别
	}()

	global.Logger.Info("日志记录器，初始化成功")
	router.InitRouter() // 初始化路由
}
