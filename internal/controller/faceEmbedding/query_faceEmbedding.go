package faceEmbedding

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/FaceEmbedding"
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
