package main

// @title 海关项目人脸识别、步态识别、高亢伪接口文档
// @version 1.0
// @description 第一个版本

// @host 172.21.116.147:8082
// @BasePath /

import (
	"fmt"
	"github.com/customs_database_server/config"
	Controllerkafka "github.com/customs_database_server/controller/kafka"
	"github.com/customs_database_server/model"
	"github.com/customs_database_server/router"
)

func main() {
	config.InitDB() // 初始化数据库
	config.InitRedis()
	fmt.Println("connect database successful...version 4")
	model.InitModel() // 初始化数据库表
	fmt.Println("build face database")
	fmt.Println("finish face database")
	go func() {
		Controllerkafka.GetImgFromKafka() // 不断从kafka接收图片，进行识别
	}()
	router.InitRouter() // 初始化路由
}
