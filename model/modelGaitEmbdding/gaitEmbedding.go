package modelGaitEmbdding

import (
	"github.com/jinzhu/gorm"
)

// ans = {"face_id":[],"name":[],"embedding":[],"face_img_url":[],"gait_img_url":[]}
type GaitEmbedding struct {
	gorm.Model `json:"-"`
	FaceId     *string `gorm:"column:face_id type:int unsigned;not null;omitempty" json:"face_id"`
	Name       *string `gorm:"column:name;type:varchar(50);not null;omitempty" json:"name"`
	Embedding  *string `gorm:"column:embedding;type:LONGTEXT;not null;omitempty" json:"embedding"`
	FaceImgURL *string `gorm:"column:face_img_url;type:varchar(255);not null;omitempty" json:"face_img_url"`
	GaitImgURL *string `gorm:"column:gait_img_url;type:varchar(255);not null;omitempty" json:"gait_img_url"`
}

//
//func CreateGait(gait *GaitEmbedding) bool {
//	create := config.DB.Create(gait)
//	if create.Error != nil {
//		panic("数据库创建步态信息错误")
//		return false
//	}
//	return true
//}
//
//func GetAllGait() []GaitEmbedding {
//	allGait := make([]GaitEmbedding, 0)
//	find := config.DB.Model(&GaitEmbedding{}).Find(&allGait)
//	if find.Error != nil {
//		panic("获取步态库失败")
//		return nil
//	}
//	return allGait
//}
