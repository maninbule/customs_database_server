package mysqlGaitResult

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelGaitEmbdding"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateGait(g *modelGaitEmbdding.GaitEmbedding) bool {
	create := config.DB.Create(g)
	if create.Error != nil {
		panic("数据库存储错误 CreateGait")
		return false
	}
	return true
}

func GaitCount() int64 {
	var cnt int64
	count := config.DB.Model(&modelGaitEmbdding.GaitEmbedding{}).Count(&cnt)
	if count.Error != nil {
		panic("数据库获取步态结果个数失败")
		return 0
	}
	return cnt
}

func GetFaceByLR(l, r int64) []modelGaitEmbdding.GaitEmbedding {
	allGait := make([]modelGaitEmbdding.GaitEmbedding, 0)
	query := config.DB.Model(&modelGaitEmbdding.GaitEmbedding{}).Offset(l - 1).Limit(r - l + 1).Find(&allGait)
	if query.Error != nil {
		panic("sql执行错误[分页查询步态识别结果失败]")
	}
	return allGait
}

func CreateQuery() *gorm.DB {
	return config.DB.Model(&modelGaitEmbdding.GaitEmbedding{})
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

func GetResult(db *gorm.DB) []modelGaitEmbdding.GaitEmbedding {
	gaits := make([]modelGaitEmbdding.GaitEmbedding, 0)
	find := db.Find(&gaits)
	if find.Error != nil {
		return nil
	}
	return gaits
}
