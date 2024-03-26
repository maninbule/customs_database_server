package faceResult

import (
	"github.com/customs_database_server/controller/response"
	mysqlFaceResult "github.com/customs_database_server/dao/mysql/FaceResult"
	"github.com/gin-gonic/gin"
)

func QueryAllFace(c *gin.Context) {
	allFace := mysqlFaceResult.GetAllFace()
	if allFace == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	response.ResponseOKWithData(c, allFace)
}
