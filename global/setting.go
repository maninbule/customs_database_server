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
)

func SetupSetting() {
	setting, err := setting.NewSetting()
	if err != nil {
		panic("viper配置文件没有加载成功")
	}
	err1 := setting.ReadSection("Server", &ServerSetting)
	err2 := setting.ReadSection("App", &AppSetting)
	err3 := setting.ReadSection("Database", &DatabaseSetting)
	if err1 != nil || err2 != nil || err3 != nil {
		panic("viper配置文件section没有读取成功")
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	fmt.Printf("%#v\n", ServerSetting)
	fmt.Printf("%#v\n", AppSetting)
	fmt.Printf("%#v\n", DatabaseSetting)
}
