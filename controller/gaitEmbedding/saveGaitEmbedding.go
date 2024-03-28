package gaitEmbedding

import (
	"github.com/customs_database_server/controller/response"
	mysqlGaitEmbedding "github.com/customs_database_server/dao/mysql/GaitEmbedding"
	"github.com/customs_database_server/model/modelGaitEmbdding"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"os"
)

type requestFormat struct {
	FaceId      string `json:"face_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Embedding   string `json:"embedding" binding:"required"`
	FaceImgUrl  string `json:"face_img_url" binding:"required"`
	GaitImgFile string `json:"gait_img_file"`
}

func SaveGaitEmbedding(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(50 << 20); err != nil {
		response.ResponseErr(c, response.CodeErrReQuestTooLarge)
		return
	}
	req := getReq(c)
	if req == nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	path := getFilePath(c)
	if path == nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "没有传输gait_img字段以及对应的文件")
		return
	}
	var modelGait modelGaitEmbdding.GaitEmbedding
	modelGait.FaceId = &req.FaceId
	modelGait.Name = &req.Name
	modelGait.Embedding = &req.Embedding
	modelGait.FaceImgURL = &req.FaceImgUrl
	modelGait.GaitImgURL = path
	if ok := mysqlGaitEmbedding.CreateGait(&modelGait); !ok {
		response.ResponseErr(c, response.CodeErrDataBase)
		err := os.Remove(*path)
		if err != nil {
			return
		}
		return
	} else {
		response.ResponseOK(c)
		return
	}
}

func getReq(c *gin.Context) *requestFormat {
	var req requestFormat
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil
	}
	return &req
}

func getFilePath(c *gin.Context) *string {
	file, err := c.FormFile("gait_img_file")
	if err != nil {
		return nil
	}
	path, err := util.SaveFileGaitDataBase(c, file)
	if err != nil {
		return nil
	}
	return &path
}
