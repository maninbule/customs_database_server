package model

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/jinzhu/gorm"
	"time"
)

type Attribute struct {
	gorm.Model
	AttrID   uint      `gorm:"column:attrId;type:int unsigned;not null"`
	Name     string    `gorm:"column:name;type:varchar(50);"`
	Hat      bool      `gorm:"column:hat;type:TINYINT(1);not null"`
	Glasses  bool      `gorm:"column:glasses;type:TINYINT(1);not null"`
	Mask     bool      `gorm:"column:mask;type:TINYINT(1);not null"`
	FaceTime time.Time `gorm:"column:faceTime;not null"`
	FaceImg  string    `gorm:"column:faceImg;type:LONGTEXT;not null"`
}

func CreateAttr(attr *Attribute) bool {
	create := config.DB.Create(attr)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllAttr() []Attribute {
	allAttr := make([]Attribute, 0)
	query := config.DB.Model(&Attribute{}).Find(&allAttr)
	if query.Error != nil {
		fmt.Println("query : ", query.Error)
		return nil
	}
	return allAttr
}

func GetAttrByTime(startDate, endDate time.Time) []Attribute {
	allAttr := make([]Attribute, 0)
	query := config.DB.Model(&Attribute{}).Where("faceTime between ? and ?", startDate, endDate).Find(&allAttr)
	if query.Error != nil {
		fmt.Println("query : ", query.Error)
		return nil
	}
	return allAttr
}
