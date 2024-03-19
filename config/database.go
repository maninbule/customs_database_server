package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	//dsn := "root:123456@(127.0.0.1:3306)/gin_info?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@(172.21.116.147:33306)/customs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic("数据库初始化失败: ")
	}
	if db.DB().Ping() != nil {
		fmt.Println(err)
		panic("数据库初始化失败2")
	}
	DB = db
}
