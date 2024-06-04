package faceResult

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/FaceResult"
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
