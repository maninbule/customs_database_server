package modelFaceResult

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 数据库用
type Face struct {
	gorm.Model     `json:"-"`
	FaceId         *uint      `gorm:"column:faceId;type:int unsigned;not null;omitempty" json:"faceId"`
	Name           *string    `gorm:"column:name;type:varchar(50);not null;omitempty" json:"name"`
	FaceTime       *time.Time `gorm:"column:faceTime;not null;omitempty" json:"faceTime"`
	CameraID       *string    `gorm:"column:cameraID;not null;omitempty" json:"cameraID"`
	FaceImgCorrect *string    `gorm:"column:faceImgCorrect;type:varchar(255);not null;omitempty" json:"faceImgCorrect"`
	FaceImgPredict *string    `gorm:"column:faceImgPredict;type:varchar(255);not null;omitempty" json:"faceImgPredict"`
	Accuracy       *float32   `gorm:"column:accuracy;type:float;not null;omitempty" json:"Accuracy"`
}

// 前端使用
type FaceJson struct {
	FaceId         string `json:"face_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	FaceTime       string `json:"face_Time" binding:"required"`
	FaceImgCorrect string `json:"face_Img_correct" binding:"required"`
	FaceImgPredict string `json:"face_Img_predict" binding:"required"`
	CameraID       string `json:"camera_id" binding:"required"`
}

// redis使用
type RedisInFaceToPredict struct {
	FaceTime string
	CameraID string
	ImgData  []byte
}

type RedisOutFaceResult struct {
	FaceId         uint
	Name           string
	FaceTime       string
	CameraID       string
	FaceImgCorrect []byte
	FaceImgPredict []byte
	Accuracy       float32
}

func Convert(r *RedisOutFaceResult) (*Face, error) {
	var res Face
	res.Name = &r.Name
	res.CameraID = &r.CameraID
	res.Accuracy = &r.Accuracy
	toTime, err := util.StringToTime(r.FaceTime)
	if err != nil {
		return nil, err
	}
	res.FaceTime = toTime
	path1, err := util.SaveFileFaceDataBaseFromByte(r.FaceImgCorrect, config.Filesuf)
	if err != nil {
		return nil, err
	}
	path2, err := util.SaveFileFaceDataBaseFromByte(r.FaceImgPredict, config.Filesuf)
	if err != nil {
		return nil, err
	}
	res.FaceImgCorrect = &path1
	res.FaceImgPredict = &path2
	return &res, nil
}

// BeforeSave hook to set FaceTime to UTC before saving to database
func (face *Face) BeforeSave() (err error) {
	if face.FaceTime != nil {
		*face.FaceTime = face.FaceTime.UTC()
	}
	return nil
}

func (f *Face) ConvertUTCtoLocalTime(location string) error {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	t := f.FaceTime.In(loc)
	f.FaceTime = &t
	return nil
}
