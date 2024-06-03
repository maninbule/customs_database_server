package responseModel

import (
	"time"
)

type Face struct {
	FaceId         *uint      `json:"faceId"`         // 人员id
	Name           *string    `json:"name"`           // 姓名
	FaceTime       *time.Time `json:"faceTime"`       // 拍摄时间
	CameraID       *string    `json:"cameraID"`       // 摄像头id
	FaceImgCorrect *string    `json:"faceImgCorrect"` // 实际数据库照片url 需加上ip:端口前缀才可以访问
	FaceImgPredict *string    `json:"faceImgPredict"` // 拍摄的照片url 需加上ip:端口前缀才可以访问
}
