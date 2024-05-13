package modelAttr

import (
	"github.com/jinzhu/gorm"
	"time"
)

// mysql
type Attribute struct {
	gorm.Model
	AttrID   *string    `gorm:"column:attr_id;type:varchar(50);not null"`
	Name     *string    `gorm:"column:name;type:varchar(50);"`
	Hat      *bool      `gorm:"column:hat;type:TINYINT(1);not null"`
	Glasses  *bool      `gorm:"column:glasses;type:TINYINT(1);not null"`
	Mask     *bool      `gorm:"column:mask;type:TINYINT(1);not null"`
	CameraId *string    `gorm:"column:camera_id;not null"`
	FaceTime *time.Time `gorm:"column:face_time;not null"`
	FaceImg  *string    `gorm:"column:face_img;type:LONGTEXT;not null"`
}

// json : 用于存储的请求参数
type AttributeJson struct {
	gorm.Model
	AttrID   string    `json:"attr_id" binding:"required" form:"attr_id"`
	Name     string    `json:"name" binding:"required" form:"name"`
	Hat      string    `json:"hat" binding:"required" form:"hat"`
	Glasses  string    `json:"glasses" binding:"required" form:"glasses"`
	Mask     string    `json:"mask" binding:"required" form:"mask"`
	CameraId string    `json:"camera_id" binding:"required" form:"camera_id"`
	FaceTime time.Time `json:"face_time" binding:"required" form:"face_time"`
	FaceImg  string    `json:"face_img"` // 存储文件之后，把路径填在这里
}

// json: 用于查询结果的请求参数
type QueryCondition struct {
	ID        string `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Mask      string `json:"mask" form:"mask"`
	Hat       string `json:"hat" form:"hat"`
	Glasses   string `json:"glasses" form:"glasses"`
	TimeStart string `json:"timeStart" form:"timeStart"`
	TimeEnd   string `json:"timeEnd" form:"timeEnd"`
	CameraID  string `json:"cameraID" form:"cameraID"`
}

func AttributeJsonToAttribute(attr *AttributeJson) *Attribute {
	var ans Attribute
	ans.AttrID = &attr.AttrID
	ans.Name = &attr.Name
	hat, glasses, mask := false, false, false
	if attr.Hat == "true" {
		hat = true
	}
	if attr.Glasses == "true" {
		glasses = true
	}
	if attr.Mask == "true" {
		mask = true
	}
	ans.Hat = &hat
	ans.Glasses = &glasses
	ans.Mask = &mask
	ans.FaceTime = &attr.FaceTime
	ans.FaceImg = &attr.FaceImg
	ans.CameraId = &attr.CameraId
	return &ans
}
