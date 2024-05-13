package gaitEmbedding

import (
	"fmt"
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
	fmt.Println("enter function SaveGaitEmbedding")
	if err := c.Request.ParseMultipartForm(50 << 20); err != nil {
		response.ResponseErr(c, response.CodeErrReQuestTooLarge)
		return
	}
	fmt.Println("解析表单通过")
	req := getReq(c)
	if req == nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	fmt.Println("绑定json成功")
	path := getFilePath(c)
	if path == nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "没有传输gait_img字段以及对应的文件")
		return
	}
	fmt.Println("图片存储成功")
	var modelGait modelGaitEmbdding.GaitEmbedding
	modelGait.FaceId = &req.FaceId
	modelGait.Name = &req.Name
	modelGait.Embedding = &req.Embedding
	modelGait.FaceImgURL = &req.FaceImgUrl
	modelGait.GaitImgURL = path
	fmt.Println(modelGait)
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
	req.FaceId = c.PostForm("face_id")
	req.Name = c.PostForm("name")
	req.Embedding = c.PostForm("embedding")
	req.FaceImgUrl = c.PostForm("face_img_url")
	if len(req.FaceId)*len(req.FaceImgUrl)*len(req.Name)*len(req.Embedding) == 0 {
		fmt.Println("绑定参数失败")
		return nil
	}
	//if err := c.ShouldBind(&req); err != nil {
	//
	//	fmt.Println("绑定参数失败")
	//	fmt.Println(err)
	//	return nil
	//}
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
