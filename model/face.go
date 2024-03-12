package model

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Face struct {
	gorm.Model
	FaceId   uint      `gorm:"column:faceId;type:int unsigned;not null"`
	Name     string    `gorm:"column:name;type:varchar(50);not null"`
	FaceImg  string    `gorm:"column:faceImg;type:LONGTEXT;not null"`
	FaceTime time.Time `gorm:"column:faceTime;not null"`
}

func CreateFace(face *Face) bool {
	//if config.DB.NewRecord(face) {
	//	return false
	//}
	create := config.DB.Create(face)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllFace() []Face {
	allFace := make([]Face, 0)
	query := config.DB.Model(&Face{}).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[获取人脸数据失败]")
	}
	return allFace
}

func GetFaceByTime(startDate, endDate time.Time) []Face {
	allFace := make([]Face, 0)
	query := config.DB.Model(&Face{}).Where("faceTime between ? and ?", startDate, endDate).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[获取指定日期人脸数据失败]")
	}
	return allFace
}

func GetFaceByID(id int) Face {
	face := Face{}
	query := config.DB.Model(&Face{}).Where("id = ?", id).First(&face)
	if query != nil {
		panic("sql执行错误[根据id查询人脸失败]")
	}
	return face
}
