package main

// @title 海关项目人脸识别、步态识别、高抗伪接口文档
// @version 1.0
// @description 第一个版本

// @host 172.21.116.147:8082
// @BasePath /

import (
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/router"
)

func main() {
	global.SetupSetting() // 读取并解析配置文件
	global.InitDBEngine() // 初始化数据库
	global.InitLogger()   // 初始化日志记录器
	//config.InitRedis()
	//fmt.Println("connect database successful...version 4")
	//model.InitModel() // 初始化数据库表
	//fmt.Println("build face database")
	//fmt.Println("finish face database")
	//go func() {
	//	Controllerkafka.GetImgFromKafka() // 不断从kafka接收图片，进行识别
	//}()
	global.Logger.Info("日志记录器，初始化成功")
	router.InitRouter() // 初始化路由
}
