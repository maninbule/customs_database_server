package responseModel

import (
	"time"
)

type Face struct {
	FaceId         *uint      `json:"faceId"`
	Name           *string    `json:"name"`
	FaceTime       *time.Time `json:"faceTime"`
	CameraID       *string    `json:"cameraID"`
	FaceImgCorrect *string    `json:"faceImgCorrect"`
	FaceImgPredict *string    `json:"faceImgPredict"`
}
