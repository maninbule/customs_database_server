package main

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model"
	"github.com/customs_database_server/router"
)

func main() {
	config.InitDB() // 初始化数据库
	config.InitRedis()
	fmt.Println("connect database successful...2222")
	model.InitModel() // 初始化数据库表
	fmt.Println("build face database")
	fmt.Println("finish face database")
	router.InitRouter() // 初始化路由

}
