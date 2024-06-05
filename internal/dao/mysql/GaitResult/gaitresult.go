package mysqlGaitResult

import (
	"fmt"
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/model/modelGait"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateGait(g *modelGaitResult.Gait) bool {
	g.BeforeSave()
	create := global.DB.Create(g)
	if create.Error != nil {
		panic("数据库存储错误 CreateGait")
		return false
	}
	return true
}

func GaitCount() int64 {
	var cnt int64
	count := global.DB.Model(&modelGaitResult.Gait{}).Count(&cnt)
	if count.Error != nil {
		panic("数据库获取步态结果个数失败")
		return 0
	}
	return cnt
}

func GetFaceByLR(l, r int64) []modelGaitResult.Gait {
	allGait := make([]modelGaitResult.Gait, 0)
	query := global.DB.Model(&modelGaitResult.Gait{}).Offset(l - 1).Limit(r - l + 1).Find(&allGait)
	if query.Error != nil {
		panic("sql执行错误[分页查询步态识别结果失败]")
	}
	return allGait
}

func CreateQuery() *gorm.DB {
	return global.DB.Model(&modelGaitResult.Gait{})
}

func GetGaitByCameraId(db *gorm.DB, id string) *gorm.DB {
	return db.Where("camera_id = ?", id)
}

func GetFaceByTimeInterval(db *gorm.DB, startTime, endTime time.Time) *gorm.DB {
	return db.Where("face_time between ? and ?", startTime, endTime)
}

func GetFaceByName(db *gorm.DB, name string) *gorm.DB {
	return db.Where("name = ?", name)
}

func GetFaceById(db *gorm.DB, id int64) *gorm.DB {
	return db.Where("face_id = ?", id)
}

func GetResult(db *gorm.DB) []modelGaitResult.Gait {
	gaits := make([]modelGaitResult.Gait, 0)
	find := db.Find(&gaits)
	if find.Error != nil {
		return nil
	}
	for i, _ := range gaits {
		err := gaits[i].ConvertUTCtoLocalTime("Asia/Shanghai")
		if err != nil {
			return nil
		}
	}
	return gaits
}
func GetResultWithLimit(db *gorm.DB, offset, limit int64) []modelGaitResult.Gait {
	gaits := make([]modelGaitResult.Gait, 0)
	db.Offset(offset).Limit(limit)
	find := db.Find(&gaits)
	if find.Error != nil {
		return nil
	}
	for i, _ := range gaits {
		err := gaits[i].ConvertUTCtoLocalTime("Asia/Shanghai")
		if err != nil {
			return nil
		}
	}
	return gaits
}

func GetCountWithCondition(db *gorm.DB) int64 {
	var count int64
	re := db.Count(&count)
	if re.Error != nil {
		fmt.Println("err = ", re.Error)
		return -1
	}
	return count
}
