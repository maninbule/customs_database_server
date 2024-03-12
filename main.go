package main

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model"
	"github.com/customs_database_server/router"
)

func main() {
	config.InitDB() // 初始化数据库
	fmt.Println("connect database successful...")
	model.InitModel()
	router.InitRouter()
}