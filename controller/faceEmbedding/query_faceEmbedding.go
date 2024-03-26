package faceEmbedding

import (
	"github.com/customs_database_server/controller/response"
	mysqlFaceEmbedding "github.com/customs_database_server/dao/mysql/FaceEmbedding"
	"github.com/gin-gonic/gin"
)

func GetAllFaceEmbedding(c *gin.Context) {
	allFace := mysqlFaceEmbedding.GetAllFace()
	if allFace == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, allFace)
		return
	}
}
