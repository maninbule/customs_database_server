package mysqlFaceResult

import (
	"fmt"
	"github.com/customs_database_server/config"
	modelFaceResult "github.com/customs_database_server/model/modelFace"
	"github.com/jinzhu/gorm"
	"sort"
	"time"
)

func CreateFace(face *modelFaceResult.Face) bool {
	//if config.DB.NewRecord(faceImgDataBase) {
	//	return false
	//}
	create := config.DB.Create(face)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllFace() []modelFaceResult.Face {
	allFace := make([]modelFaceResult.Face, 0)
	query := config.DB.Model(&modelFaceResult.Face{}).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[获取人脸数据失败]")
	}
	return allFace
}

func GetFaceByTime(startDate, endDate time.Time) []modelFaceResult.Face {
	allFace := make([]modelFaceResult.Face, 0)
	query := config.DB.Model(&modelFaceResult.Face{}).Where("faceTime between ? and ?", startDate, endDate).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[获取指定日期人脸数据失败]")
	}
	return allFace
}

func GetFaceByID(id int) modelFaceResult.Face {
	face := modelFaceResult.Face{}
	query := config.DB.Model(&modelFaceResult.Face{}).Where("id = ?", id).First(&face)
	if query != nil {
		panic("sql执行错误[根据id查询人脸失败]")
	}
	return face
}

func GetFaceByLR(l, r int64) []modelFaceResult.Face {
	allFace := make([]modelFaceResult.Face, 0)
	query := config.DB.Model(&modelFaceResult.Face{}).Offset(l - 1).Limit(r - l + 1).Find(&allFace)
	if query.Error != nil {
		panic("sql执行错误[根据id区间查询人脸失败]")
	}
	return allFace
}

func GetCount() int64 {
	var cnt int64
	query := config.DB.Model(&modelFaceResult.Face{}).Count(&cnt)
	if query.Error != nil {
		panic("sql执行错误[查询count失败]")
	}
	return cnt
}

func CreateQuery() *gorm.DB {
	return config.DB.Model(&modelFaceResult.Face{})
}

func GetFaceByCameraID(db *gorm.DB, id string) *gorm.DB {
	return db.Where("cameraID = ?", id)
}

func GetFaceByTimeInterval(db *gorm.DB, startTime, endTime time.Time) *gorm.DB {
	fmt.Println(startTime, endTime)
	return db.Where("faceTime between ? and ?", startTime, endTime)
}

func GetFaceByName(db *gorm.DB, name string) *gorm.DB {
	return db.Where("name = ?", name)
}

func GetFaceById(db *gorm.DB, id int64) *gorm.DB {
	return db.Where("faceId = ?", id)
}

func GetResult(db *gorm.DB) []modelFaceResult.Face {
	faces := make([]modelFaceResult.Face, 0)
	find := db.Find(&faces)
	if find.Error != nil {
		fmt.Println("find.Error = ", find.Error)
		return nil
	}
	sort.Slice(faces, func(i, j int) bool {
		return faces[i].FaceTime.After(*faces[j].FaceTime)
	})
	return faces
}
