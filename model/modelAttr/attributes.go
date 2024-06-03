package modelAttr

import (
	"github.com/jinzhu/gorm"
	"time"
)

// mysql
type Attribute struct {
	gorm.Model `json:"-"`
	AttrID     *string    `gorm:"column:attr_id;type:varchar(50);not null"` // id字段
	Name       *string    `gorm:"column:name;type:varchar(50);"`            // 人员名字 保留字段，目前不使用
	Hat        *bool      `gorm:"column:hat;type:TINYINT(1);not null"`      // 是否佩戴帽子 true 或者 false
	Glasses    *bool      `gorm:"column:glasses;type:TINYINT(1);not null"`  // 是否佩戴眼镜 true 或者 false
	Mask       *bool      `gorm:"column:mask;type:TINYINT(1);not null"`     // 是否佩戴口罩 true 或者 false
	CameraId   *string    `gorm:"column:camera_id;not null"`                // 摄像头id
	FaceTime   *time.Time `gorm:"column:face_time;not null"`                // 拍摄时间
	FaceImg    *string    `gorm:"column:face_img;type:LONGTEXT;not null"`   // 图片url 需要加上ip:端口前缀才能访问
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
	ID        string `json:"id" form:"id" example:"123" `                              // ID 是一个用于查询的唯一标识符
	Name      string `json:"name" form:"name" example:"小明"`                            // Name 是一个用于查询的姓名
	Mask      string `json:"mask" form:"mask" example:"true"`                          // Mask 表示是否戴口罩，可选值为 "true" 或 "false"
	Hat       string `json:"hat" form:"hat" example:"false"`                           // Hat 表示是否戴帽子，可选值为 "true" 或 "false"
	Glasses   string `json:"glasses" form:"glasses" example:"true"`                    // Glasses 表示是否戴眼镜，可选值为 "true" 或 "false"
	TimeStart string `json:"timeStart" form:"timeStart" example:"2024-05-17 10:00:00"` // TimeStart 表示查询开始时间
	TimeEnd   string `json:"timeEnd" form:"timeEnd" example:"2024-05-17 18:00:00"`     // TimeEnd 表示查询结束时间
	CameraID  string `json:"cameraID" form:"cameraID" example:"摄像头1"`                  // CameraID 表示摄像头的唯一标识符
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

func (attr *Attribute) BeforeSave() (err error) {
	if attr.FaceTime != nil {
		*attr.FaceTime = attr.FaceTime.UTC()
	}
	return nil
}

func (attr *Attribute) ConvertUTCtoLocalTime(location string) error {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	t := attr.FaceTime.In(loc)
	attr.FaceTime = &t
	return nil
}
