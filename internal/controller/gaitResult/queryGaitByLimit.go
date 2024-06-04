package gaitResult

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/GaitResult"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 选取查询结果中的第l到第r条记录
func QueryFaceByLimit(c *gin.Context) {
	left := c.Param("left")
	right := c.Param("right")
	l, err1 := strconv.ParseInt(left, 0, 0)
	r, err2 := strconv.ParseInt(right, 0, 0)
	total := mysqlGaitResult.GaitCount()
	if err1 != nil || err2 != nil || l > r || l < 1 || r > total {
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	allGait := mysqlGaitResult.GetFaceByLR(l, r)
	if allGait == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, allGait)
	}

}
