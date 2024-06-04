package faceResult

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/FaceResult"
	"github.com/gin-gonic/gin"
)

func GetFaceCount(c *gin.Context) {
	cnt := mysqlFaceResult.GetCount()
	response.ResponseOKWithData(c, cnt)
}
