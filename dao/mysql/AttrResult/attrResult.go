package mysqlAttrResult

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelAttr"
	"github.com/jinzhu/gorm"
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

func GetResult(db *gorm.DB) []modelAttr.Attribute {
	gaits := make([]modelAttr.Attribute, 0)
	find := db.Find(&gaits)
	if find.Error != nil {
		return nil
	}
	return gaits
}
