package faceEmbedding

import (
	"errors"
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelFaceEemdding"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type requestFormat struct {
	FaceId    string `json:"face_id"`
	Name      string `json:"name"`
	Embedding string `json:"embedding"`
	FaceImg   string `json:"face_img"`
}

func (r requestFormat) Valid() bool {
	keys := []string{
		r.Name, r.FaceId, r.FaceImg, r.FaceImg,
	}
	for _, key := range keys {
		if len(key) == 0 {
			return false
		}
	}
	if _, err := r.GetFaceId(); err != nil {
		return false
	}
	return true
}

func (r requestFormat) GetFaceId() (uint, error) {
	parseUint, err := strconv.ParseUint(r.FaceId, 0, 0)
	return uint(parseUint), err
}

func SaveFaceEmbedding(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseBadRequest(c, "post数据量过大")
		return
	}
	faceImg, err := c.FormFile("face_img")
	if err != nil {
		response.ResponseBadRequest(c, "传输的数据不存在对应face_img的文件")
		return
	}
	path, err := util.SaveFileFaceDataBase(c, faceImg)
	if err != nil {
		response.ResponseBadRequest(c, "图像存储到本地磁盘错误")
		return
	}
	var format requestFormat
	format.FaceId = c.PostForm("face_id")
	format.Name = c.PostForm("name")
	format.Embedding = c.PostForm("embedding")
	format.FaceImg = path
	err = saveToDB(c, &format)
	if err != nil {
		response.ResponseBadRequest(c, "数据库存储错误")
		err := os.Remove(path)
		if err != nil {
			return
		}
		return
	}
}

func saveToDB(c *gin.Context, format *requestFormat) error {
	var faceEmbedding modelFaceEemdding.FaceEmbedding
	id, _ := format.GetFaceId()
	faceEmbedding.FaceId = &id
	faceEmbedding.Name = &format.Name
	faceEmbedding.Embedding = &format.Embedding
	faceEmbedding.FaceImg = &format.FaceImg

	if ok := modelFaceEemdding.CreateFace(&faceEmbedding); !ok {
		response.ResponseBadRequest(c, "数据库存储错误")
		return errors.New("err")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "存入数据库成功",
	})
	return nil
}
