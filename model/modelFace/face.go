package modelFace

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Face struct {
	gorm.Model
	FaceId         *uint      `gorm:"column:faceId;type:int unsigned;not null;omitempty"`
	Name           *string    `gorm:"column:name;type:varchar(50);not null;omitempty"`
	FaceTime       *time.Time `gorm:"column:faceTime;not null;omitempty"`
	CameraID       *string    `gorm:"column:cameraID;not null;omitempty"`
	FaceImgCorrect *string    `gorm:"column:faceImgCorrect;type:LONGTEXT;not null;omitempty"`
	FaceImgPredict *string    `gorm:"column:faceImgPredict;type:LONGTEXT;not null;omitempty"`
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

func GetFaceByLR(l, r int64) []Face {
	allFace := make([]Face, 0)
	query := config.DB.Model(&Face{}).Offset(l - 1).Limit(r - l + 1).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[根据id区间查询人脸失败]")
	}
	return allFace
}

func GetCount() int64 {
	var cnt int64
	query := config.DB.Model(&Face{}).Count(&cnt)
	if query.Error != nil {
		panic("sql执行错误[查询count失败]")
	}
	return cnt
}

func CreateQuery() *gorm.DB {
	return config.DB.Model(&Face{})
}

func GetFaceByCameraID(db *gorm.DB, id string) *gorm.DB {
	return db.Where("cameraID = ?", id)
}

func GetFaceByTimeInterval(db *gorm.DB, startTime, endTime time.Time) *gorm.DB {
	return db.Where("faceTime between ? and ?", startTime, endTime)
}

func GetFaceByName(db *gorm.DB, name string) *gorm.DB {
	return db.Where("name = ?", name)
}

func GetFaceById(db *gorm.DB, id int64) *gorm.DB {
	return db.Where("faceId = ?", id)
}

func GetResult(db *gorm.DB) []Face {
	faces := make([]Face, 0)
	find := db.Find(&faces)
	if find.Error != nil {
		return nil
	}
	return faces
}
