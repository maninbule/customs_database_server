package face

import (
	"github.com/customs_database_server/model/modelFace"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 选取查询结果中的第l到第r条记录
func QueryFaceByLimit(c *gin.Context) {
	left := c.Param("left")
	right := c.Param("right")
	l, err1 := strconv.ParseInt(left, 0, 0)
	r, err2 := strconv.ParseInt(right, 0, 0)
	total := modelFace.GetCount()
	if err1 != nil || err2 != nil || l > r || l < 1 || r > total {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  "参数错误",
		})
		return
	}
	allFace := modelFace.GetFaceByLR(l, r)
	if allFace == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"err":  "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"data-length": len(allFace),
		"data":        allFace,
	})
}
