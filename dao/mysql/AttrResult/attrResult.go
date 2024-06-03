package mysqlAttrResult

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelAttr"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateAttr(attr *modelAttr.Attribute) bool {
	attr.BeforeSave()
	create := config.DB.Create(attr)
	if create.Error != nil {
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

func CreateQuery() *gorm.DB {
	return config.DB.Model(&modelAttr.Attribute{})
}

func GetByCameraId(db *gorm.DB, id string) *gorm.DB {
	return db.Where("camera_id = ?", id)
}

func GetByTimeInterval(db *gorm.DB, startTime, endTime time.Time) *gorm.DB {
	return db.Where("face_time between ? and ?", startTime, endTime)
}

func GetByMask(db *gorm.DB, op int) *gorm.DB {
	return db.Where("mask = ?", op)
}
func GetByHat(db *gorm.DB, op int) *gorm.DB {
	return db.Where("hat = ?", op)
}

func GetByGlasses(db *gorm.DB, op int) *gorm.DB {
	return db.Where("glasses = ?", op)
}

func GetByName(db *gorm.DB, name string) *gorm.DB {
	return db.Where("name = ?", name)
}

func GetById(db *gorm.DB, id int64) *gorm.DB {
	return db.Where("face_id = ?", id)
}

func GetResultWithLimit(db *gorm.DB, offset, limit int64) []modelAttr.Attribute {
	attrs := make([]modelAttr.Attribute, 0)
	find := db.Offset(offset).Limit(limit).Find(&attrs)
	if find.Error != nil {
		fmt.Println("Error = ", find.Error)
		return nil
	}
	for i, _ := range attrs {
		err := attrs[i].ConvertUTCtoLocalTime("Asia/Shanghai")
		if err != nil {
			fmt.Println("attrs[i].ConvertUTCtoLocalTime err = ", err)
			return nil
		}
	}
	return attrs
}

func GetResultCount(db *gorm.DB) int64 {
	var ans int64
	count := db.Count(&ans)
	if count.Error != nil {
		fmt.Println("Error = ", count.Error)
		return -1
	}
	return ans
}
