package modelFaceEemdding

import (
	"encoding/base64"
	"github.com/customs_database_server/util"
	"github.com/jinzhu/gorm"
)

// mysql
type FaceEmbedding struct {
	gorm.Model
	FaceId     *uint   `gorm:"column:faceId;type:int unsigned;not null;omitempty" json:"face_id"`
	Name       *string `gorm:"column:name;type:varchar(50);not null;omitempty" json:"name"`
	Embedding  *string `gorm:"column:embedding;type:LONGTEXT;not null;omitempty" json:"embedding"`
	FaceImgURL *string `gorm:"column:face_img_url;type:varchar(255);not null;omitempty" json:"face_img_url"`
}

// redis使用
type RedisInFaceEb struct {
	FaceId   uint
	Name     string
	fileName string
	ImgData  []byte
}

type RedisOutFaceEb struct {
	FaceId       uint
	Name         string
	ImgData      []byte
	ImgEmbedding []byte
}

// model互相转换
func RedisOutToFaceEmbedding(r *RedisOutFaceEb) (*FaceEmbedding, error) {
	var res FaceEmbedding
	res.Name = &r.Name
	res.FaceId = &r.FaceId
	toString := base64.StdEncoding.EncodeToString(r.ImgEmbedding)
	res.Embedding = &toString
	fileUrl, err := util.SaveFileFaceDataBaseFromByte(r.ImgData, ".png")
	if err != nil {
		return nil, err
	}
	res.FaceImgURL = &fileUrl
	return &res, nil
}
