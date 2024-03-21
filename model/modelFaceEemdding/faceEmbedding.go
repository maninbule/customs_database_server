package modelFaceEemdding

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/jinzhu/gorm"
)

type FaceEmbedding struct {
	gorm.Model
	FaceId     *uint   `gorm:"column:faceId;type:int unsigned;not null;omitempty" json:"face_id"`
	Name       *string `gorm:"column:name;type:varchar(50);not null;omitempty" json:"name"`
	Embedding  *string `gorm:"column:embedding;type:LONGTEXT;not null;omitempty" json:"embedding"`
	FaceImgURL *string `gorm:"column:face_img_url;type:varchar(255);not null;omitempty" json:"faceImgURL"`
}

func CreateFace(face *FaceEmbedding) bool {
	create := config.DB.Create(face)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllFace() []FaceEmbedding {
	allFace := make([]FaceEmbedding, 0)
	find := config.DB.Model(&FaceEmbedding{}).Find(&allFace)
	if find.Error != nil {
		panic("sql执行错误[获取人脸向量数据失败]： GetallFace")
		return nil
	}
	return allFace
}
