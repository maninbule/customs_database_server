package modelFaceEemdding

import (
	"encoding/base64"
	"github.com/customs_database_server/util"
	"github.com/jinzhu/gorm"
	"strconv"
)

// mysql
type FaceEmbedding struct {
	gorm.Model
	FaceId     *uint   `gorm:"column:faceId;type:int unsigned;not null;omitempty" json:"face_id"`
	Name       *string `gorm:"column:name;type:varchar(50);not null;omitempty" json:"name"`
	Embedding  *string `gorm:"column:embedding;type:LONGTEXT;not null;omitempty" json:"embedding"`
	FaceImgURL *string `gorm:"column:face_img_url;type:varchar(255);not null;omitempty" json:"face_img_url"`
}

// 前端使用: 数据库存储需要的结构体
// 前端传输的是一个formdata表单，其中face_img是一个文件的字段
// FaceImg最终要存储这个文件的路径，(此文件要先存储)
type RequestFormat struct {
	FaceId    string `json:"face_id" binding:"required" form:"face_id"`
	Name      string `json:"name" binding:"required" form:"name"`
	Embedding string `json:"embedding" binding:"required" form:"embedding"`
	FaceImg   string `json:"face_img"` // 文件，需存储文件，再填入文件路径
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

// RequestFormat 转 FaceEmbedding
func RequestFormatToFaceEmbedding(req *RequestFormat) *FaceEmbedding {
	var result FaceEmbedding
	parseUint, _ := strconv.ParseUint(req.FaceId, 0, 0)
	x := uint(parseUint)
	result.FaceId = &x
	result.Name = &req.Name
	result.Embedding = &req.Embedding
	result.FaceImgURL = &req.FaceImg
	return &result
}
