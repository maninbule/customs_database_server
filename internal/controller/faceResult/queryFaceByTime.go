package faceResult

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/FaceResult"
	"github.com/gin-gonic/gin"
	"time"
)

func QueryFaceByTime(c *gin.Context) {
	startTime := c.Param("startTime")
	endTime := c.Param("endTime")

	start, err1 := time.Parse("2006-01-02 15:04:05", startTime)
	end, err2 := time.Parse("2006-01-02 15:04:05", endTime)
	if err1 != nil || err2 != nil || end.Sub(start) < 0 {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间错误格式错误")
		return
	}

	allFace := mysqlFaceResult.GetFaceByTime(start, end)
	if allFace == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	response.ResponseOKWithData(c, allFace)
}
