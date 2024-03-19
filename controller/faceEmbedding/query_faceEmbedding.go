package faceEmbedding

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelFaceEemdding"
	"github.com/gin-gonic/gin"
)

func GetAllFaceEmbedding(c *gin.Context) {
	allFace := modelFaceEemdding.GetAllFace()
	if allFace == nil {
		response.ResponseBadRequest(c, "查询失败，服务器错误")
		return
	} else {
		response.ResponseOKWithData(c, allFace)
		return
	}
}
