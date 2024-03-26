package faceResult

import (
	"github.com/customs_database_server/controller/response"
	mysqlFaceResult "github.com/customs_database_server/dao/mysql/FaceResult"
	"github.com/gin-gonic/gin"
)

func GetFaceCount(c *gin.Context) {
	cnt := mysqlFaceResult.GetCount()
	response.ResponseOKWithData(c, cnt)
}
