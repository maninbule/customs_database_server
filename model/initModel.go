package model

import "github.com/customs_database_server/config"

func InitModel() {
	config.DB.AutoMigrate(&Face{})
	config.DB.AutoMigrate(&Attribute{})
}
