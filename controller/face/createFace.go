package face

import (
	"fmt"
	"github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
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

func CreateFace(c *gin.Context) {
	face := modelFace.Face{}
	jsonFormat := reuestFormat{}
	if ok := util.ParseBody(c, &jsonFormat); !ok || !jsonFormat.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  "传入json的格式不正确或者字段为空",
		})
		return
	}
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
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "存入数据库成功",
	})
}
