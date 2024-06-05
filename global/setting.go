package global

import (
	"fmt"
	"github.com/customs_database_server/pkg/setting"
	"time"
)

var (
	ServerSetting   setting.ServerSettings
	AppSetting      setting.AppSettings
	DatabaseSetting setting.DatabaseSettings
	RedisSetting    setting.RedisSetting
	KafkaSetting    setting.KafkaSetting
)

func SetupSetting() {
	setting, err := setting.NewSetting()
	if err != nil {
		panic("viper配置文件没有加载成功")
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		panic("viper配置文件Server section没有读取成功")
	}
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		panic("viper配置文件App section没有读取成功")
	}
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		panic("viper配置文件Database section没有读取成功")
	}
	err = setting.ReadSection("Redis", &RedisSetting)
	if err != nil {
		panic("viper配置文件Redis section没有读取成功")
	}
	err = setting.ReadSection("Kafka", &KafkaSetting)
	if err != nil {
		panic("viper配置文件Kafka section没有读取成功")
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	fmt.Printf("%#v\n", ServerSetting)
	fmt.Printf("%#v\n", AppSetting)
	fmt.Printf("%#v\n", DatabaseSetting)
	fmt.Printf("%#v\n", RedisSetting)
	fmt.Printf("%#v\n", KafkaSetting)

}
