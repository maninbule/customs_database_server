package modelGaitEmbdding

import (
	"github.com/jinzhu/gorm"
)

type GaitEmbedding struct {
	gorm.Model
	FaceId     *string `gorm:"column:face_id type:int unsigned;not null;omitempty"`
	Name       *string `gorm:"column:name;type:varchar(50);not null;omitempty"`
	Embedding  *string `gorm:"column:embedding;type:LONGTEXT;not null;omitempty"`
	FaceImgURL *string `gorm:"column:face_img_url;type:varchar(255);not null;omitempty"`
	GaitImgURL *string `gorm:"column:gait_img_url;type:varchar(255);not null;omitempty"`
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
