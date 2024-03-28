package faceResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlFaceResult "github.com/customs_database_server/dao/mysql/FaceResult"
	modelFaceResult "github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type requestFormat struct {
	FaceId         string `json:"face_id"`
	Name           string `json:"name"`
	FaceTime       string `json:"face_Time"`
	FaceImgCorrect string `json:"face_Img_correct"`
	FaceImgPredict string `json:"face_Img_predict"`
	CameraID       string `json:"camera_id"`
}

func (r requestFormat) Valid() bool {
	fmt.Println(r)
	keys := []string{
		r.FaceId, r.Name, r.FaceTime, r.FaceImgCorrect, r.FaceImgPredict, r.CameraID,
	}
	for _, key := range keys {
		if len(key) == 0 {
			return false
		}
	}
	if _, err := r.getTime(); err != nil {
		fmt.Println(err)
		return false
	}
	if _, err := r.getId(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (r requestFormat) getId() (uint, error) {
	parseUint, err := strconv.ParseUint(r.FaceId, 0, 0)
	return uint(parseUint), err
}

func (r requestFormat) getTime() (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", r.FaceTime)
	return t, err
}

func SaveFaceCompare(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseErr(c, response.CodeErrReQuestTooLarge)
		return
	}
	fileCorrectImg, ok1 := c.FormFile("face_Img_correct")
	filePredictImg, ok2 := c.FormFile("face_Img_predict")
	if ok1 != nil || ok2 != nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "缺少人脸图像字段")
		return
	}
	pathCorrectImg, err1 := util.SaveFileRecResult(c, fileCorrectImg)
	pathPredictImg, err2 := util.SaveFileRecResult(c, filePredictImg)
	if err1 != nil || err2 != nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	var jsonFormat requestFormat
	jsonFormat.FaceId = c.PostForm("face_id")
	jsonFormat.Name = c.PostForm("name")
	jsonFormat.FaceTime = c.PostForm("face_Time")
	jsonFormat.CameraID = c.PostForm("camera_id")
	jsonFormat.FaceImgCorrect = pathCorrectImg
	jsonFormat.FaceImgPredict = pathPredictImg
	if !jsonFormat.Valid() {
		response.ResponseErr(c, response.CodeErrRequestParamNotExisted)
		return
	}
	ok := CreateFace(c, &jsonFormat)
	if !ok {
		response.ResponseErr(c, response.CodeErrDataBase)
		err := os.Remove(pathCorrectImg)
		if err != nil {
			return
		}
		err2 := os.Remove(pathPredictImg)
		if err2 != nil {
			return
		}
		return
	}
	response.ResponseOK(c)
}

func CreateFace(c *gin.Context, jsonFormat *requestFormat) bool {
	face := modelFaceResult.Face{}
	id, _ := jsonFormat.getId()
	face.FaceId = &id
	face.Name = &jsonFormat.Name
	face.FaceImgCorrect = &jsonFormat.FaceImgCorrect
	face.FaceImgPredict = &jsonFormat.FaceImgPredict
	face.CameraID = &jsonFormat.CameraID
	t, _ := jsonFormat.getTime()
	addTime := t.Add(-8 * time.Hour)
	face.FaceTime = &addTime
	ok := mysqlFaceResult.CreateFace(&face)
	return ok
}
