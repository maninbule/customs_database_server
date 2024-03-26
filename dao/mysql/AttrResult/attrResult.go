package mysqlAttrResult

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelAttr"
	"time"
)

func CreateAttr(attr *modelAttr.Attribute) bool {
	create := config.DB.Create(attr)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllAttr() []modelAttr.Attribute {
	allAttr := make([]modelAttr.Attribute, 0)
	query := config.DB.Model(&modelAttr.Attribute{}).Find(&allAttr)
	if query.Error != nil {
		fmt.Println("query : ", query.Error)
		return nil
	}
	return allAttr
}

func GetAttrByTime(startDate, endDate time.Time) []modelAttr.Attribute {
	allAttr := make([]modelAttr.Attribute, 0)
	query := config.DB.Model(&modelAttr.Attribute{}).Where("faceTime between ? and ?", startDate, endDate).Find(&allAttr)
	if query.Error != nil {
		fmt.Println("query : ", query.Error)
		return nil
	}
	return allAttr
}
