package face

import (
	"errors"
	"fmt"
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type reuestFormat struct {
	FaceId         string `json:"face_id"`
	Name           string `json:"name"`
	FaceTime       string `json:"face_Time"`
	FaceImgCorrect string `json:"face_Img_correct"`
	FaceImgPredict string `json:"face_Img_predict"`
	CameraID       string `json:"camera_id"`
}

func (r reuestFormat) Valid() bool {
	keys := []string{
		r.FaceId, r.Name, r.FaceTime, r.FaceImgCorrect, r.FaceImgPredict, r.CameraID,
	}
	for _, key := range keys {
		if len(key) == 0 {
			return false
		}
	}
	if _, err := r.getTime(); err != nil {
		return false
	}
	if _, err := r.getId(); err != nil {
		return false
	}
	return true
}

func (r reuestFormat) getId() (uint, error) {
	parseUint, err := strconv.ParseUint(r.FaceId, 0, 0)
	return uint(parseUint), err
}

func (r reuestFormat) getTime() (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", r.FaceTime)
	return t, err
}

func SaveFaceCompare(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseBadRequest(c)
		return
	}
	fileCorrectImg, ok1 := c.FormFile("face_Img_correct")
	filePredictImg, ok2 := c.FormFile("face_Img_predict")
	if ok1 != nil || ok2 != nil {
		response.ResponseBadRequest(c)
		return
	}
	pathCorrectImg, err1 := util.SaveFile(c, fileCorrectImg)
	pathPredictImg, err2 := util.SaveFile(c, filePredictImg)
	if err1 != nil || err2 != nil {
		fmt.Println("存储图片错误, ", err1, err2)
		response.ResponseBadRequest(c)
		return
	}
	var jsonFormat reuestFormat
	jsonFormat.FaceId = c.PostForm("face_id")
	jsonFormat.Name = c.PostForm("name")
	jsonFormat.FaceTime = c.PostForm("face_Time")
	jsonFormat.CameraID = c.PostForm("camera_id")
	jsonFormat.FaceImgCorrect = pathCorrectImg
	jsonFormat.FaceImgPredict = pathPredictImg
	if !jsonFormat.Valid() {
		response.ResponseBadRequest(c)
		return
	}
	err = CreateFace(c, &jsonFormat)
	if err != nil {
		err := os.Remove(pathCorrectImg)
		if err != nil {
			return
		}
		err2 := os.Remove(pathPredictImg)
		if err2 != nil {
			return
		}
	}
}

func CreateFace(c *gin.Context, jsonFormat *reuestFormat) error {
	face := modelFace.Face{}
	id, _ := jsonFormat.getId()
	face.FaceId = &id
	face.Name = &jsonFormat.Name
	face.FaceImgCorrect = &jsonFormat.FaceImgCorrect
	face.FaceImgPredict = &jsonFormat.FaceImgPredict
	face.CameraID = &jsonFormat.CameraID
	t, _ := jsonFormat.getTime()
	addTime := t.Add(-8 * time.Hour)
	face.FaceTime = &addTime
	fmt.Println("face格式化之后： ", face)
	ok := modelFace.CreateFace(&face)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusOK,
			"err":  "入库失败",
		})
		return errors.New("err")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "存入数据库成功",
	})
	return nil
}
