package gaitResult

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelGait"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 选取查询结果中的第l到第r条记录
func QueryFaceByLimit(c *gin.Context) {
	left := c.Param("left")
	right := c.Param("right")
	l, err1 := strconv.ParseInt(left, 0, 0)
	r, err2 := strconv.ParseInt(right, 0, 0)
	total := modelGait.GaitCount()
	if err1 != nil || err2 != nil || l > r || l < 1 || r > total {
		response.ResponseBadRequest(c, "分页区间不合法")
		return
	}
	allGait := modelGait.GetFaceByLR(l, r)
	if allGait == nil {
		response.ResponseInternalErr(c, "从数据库获取数据失败")
		return
	} else {
		response.ResponseOKWithData(c, allGait)
	}

}
