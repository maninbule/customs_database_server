package modelGaitResult

import (
	"encoding/base64"
	"fmt"
	"github.com/customs_database_server/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Gait struct {
	gorm.Model
	FaceId     *string    `gorm:"column:face_id type:int unsigned;not null;omitempty"`
	Name       *string    `gorm:"column:name;type:varchar(50);not null;omitempty"`
	CameraID   *string    `gorm:"column:camera_id;not null;omitempty"`
	FaceTime   *time.Time `gorm:"column:face_time;not null;omitempty"`
	FaceImgURL *string    `gorm:"column:face_img_url;type:varchar(255);not null;omitempty"`
	GaitImgURL *string    `gorm:"column:gait_img_url;type:varchar(255);not null;omitempty"`
}

// 请求参数

type GaitRequest struct {
	FaceId     string `json:"face_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	CameraID   string `json:"camera_id" binding:"required"`
	FaceTime   string `json:"face_time" binding:"required"`
	FaceImgURL string `json:"face_img_url" binding:"required"`
	GaitImg    string `json:"gait_img" binding:"required"`
}

func ConvertGaitRequestToGait(g *GaitRequest) (*Gait, bool) {
	var res Gait
	res.FaceId = &g.FaceId
	res.Name = &g.Name
	res.CameraID = &g.CameraID
	var t time.Time
	ok := util.ParseTime(g.FaceTime, &t)
	if !ok {
		return nil, false
	}
	res.FaceTime = &t
	res.FaceImgURL = &g.FaceImgURL
	bytes, err2 := base64.StdEncoding.DecodeString(g.GaitImg)
	if err2 != nil {
		return nil, false
	}
	path, err := util.SaveFileFaceDataBaseFromByte(bytes, ".png")
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	res.GaitImgURL = &path
	return &res, true
}

//
//func CreateGait(g *Gait) bool {
//	create := config.DB.Create(g)
//	if create.Error != nil {
//		panic("数据库存储错误 CreateGait")
//		return false
//	}
//	return true
//}
//
//func GaitCount() int64 {
//	var cnt int64
//	count := config.DB.Model(&Gait{}).Count(&cnt)
//	if count.Error != nil {
//		panic("数据库获取步态结果个数失败")
//		return 0
//	}
//	return cnt
//}
//
//func GetFaceByLR(l, r int64) []Gait {
//	allGait := make([]Gait, 0)
//	query := config.DB.Model(&Gait{}).Offset(l - 1).Limit(r - l + 1).Find(&allGait)
//	if query.Error != nil {
//		panic("sql执行错误[分页查询步态识别结果失败]")
//	}
//	return allGait
//}
//
//func CreateQuery() *gorm.DB {
//	return config.DB.Model(&Gait{})
//}
//
//func GetGaitByCameraId(db *gorm.DB, id string) *gorm.DB {
//	return db.Where("camera_id = ?", id)
//}
//
//func GetFaceByTimeInterval(db *gorm.DB, startTime, endTime time.Time) *gorm.DB {
//	return db.Where("face_time between ? and ?", startTime, endTime)
//}
//
//func GetFaceByName(db *gorm.DB, name string) *gorm.DB {
//	return db.Where("name = ?", name)
//}
//
//func GetFaceById(db *gorm.DB, id int64) *gorm.DB {
//	return db.Where("face_id = ?", id)
//}
//
//func GetResult(db *gorm.DB) []Gait {
//	gaits := make([]Gait, 0)
//	find := db.Find(&gaits)
//	if find.Error != nil {
//		return nil
//	}
//	return gaits
//}
