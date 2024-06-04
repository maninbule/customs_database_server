package faceEmbedding

import (
	"fmt"
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/FaceEmbedding"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func paramRequest(c *gin.Context) (*modelFaceEemdding.RequestFormat, error) {
	// 将formdata的内容存入requestFormat
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseErr(c, response.CodeErrReQuestTooLarge)
		return nil, err
	}
	// 绑定参数
	var req modelFaceEemdding.RequestFormat
	if err = c.ShouldBind(&req); err != nil {
		return nil, err
	}
	// 检查face_id是否是整数
	_, err = strconv.ParseUint(req.FaceId, 0, 0)
	if err != nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "face_id字段非法")
		return nil, err
	}
	// 存储文件
	faceImg, err := c.FormFile("face_img")
	if err != nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "传输的数据不存在对应face_img的文件")
		return nil, err
	}
	path, err := util.SaveFileFaceDataBase(c, faceImg)
	if err != nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return nil, err
	}
	req.FaceImg = path
	return &req, nil
}

func SaveFaceEmbedding(c *gin.Context) {
	fmt.Println("enter SaveFaceEmbedding")
	req, err := paramRequest(c)
	if err != nil {
		fmt.Println("SaveFaceEmbedding : ", err)
		return
	}
	embedding := modelFaceEemdding.RequestFormatToFaceEmbedding(req)
	ok := mysqlFaceEmbedding.CreateFace(embedding)
	if !ok { // 存储不成功
		fmt.Println("!ok 存储不成功")
		response.ResponseErr(c, response.CodeErrDataBase)
		err := os.Remove(req.FaceImg)
		if err != nil {
			//response.ResponseErr(c, response.CodeErrDataBase)
			return
		}
		return
	} else {
		response.ResponseOKWithData(c, gin.H{"img_url": embedding.FaceImgURL})
	}
}
