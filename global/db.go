package global

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func InitDBEngine() {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		DatabaseSetting.Username,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.DBName,
		DatabaseSetting.Charset,
		DatabaseSetting.ParseTime)
	db, err := gorm.Open(DatabaseSetting.DBType, s)
	if err != nil {
		fmt.Printf(err.Error())
		panic("数据库连接配置错误")
	}
	if ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	//db.SingularTable(true)
	db.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	DB = db
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Host,
		Password: RedisSetting.Password,
		DB:       RedisSetting.DB,
	})
	if Redis == nil {
		panic("Redis初始化失败")
	}
}
