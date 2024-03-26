package modelAttr

import (
	"github.com/jinzhu/gorm"
	"time"
)

// mysql
type Attribute struct {
	gorm.Model
	AttrID   *string    `gorm:"column:attrId;type:varchar(50);not null"`
	Name     *string    `gorm:"column:name;type:varchar(50);"`
	Hat      *bool      `gorm:"column:hat;type:TINYINT(1);not null"`
	Glasses  *bool      `gorm:"column:glasses;type:TINYINT(1);not null"`
	Mask     *bool      `gorm:"column:mask;type:TINYINT(1);not null"`
	FaceTime *time.Time `gorm:"column:faceTime;not null"`
	FaceImg  *string    `gorm:"column:faceImg;type:LONGTEXT;not null"`
}

// json
type AttributeJson struct {
	gorm.Model
	AttrID   string    `json:"attr_id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Hat      bool      `json:"hat" binding:"required"`
	Glasses  bool      `json:"glasses" binding:"required"`
	Mask     bool      `json:"mask" binding:"required"`
	FaceTime time.Time `json:"face_time" binding:"required"`
	FaceImg  string    `json:"face_img" binding:"required"`
}
