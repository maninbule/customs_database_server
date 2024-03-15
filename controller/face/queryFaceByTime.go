package face

import (
	"fmt"
	"github.com/customs_database_server/model/modelFace"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func QueryFaceByTime(c *gin.Context) {
	startTime := c.Param("startTime")
	endTime := c.Param("endTime")

	start, err1 := time.Parse("2006-01-02 15:04:05", startTime)
	end, err2 := time.Parse("2006-01-02 15:04:05", endTime)
	if err1 != nil || err2 != nil || end.Sub(start) < 0 {
		fmt.Println("time.Parse ", err1, err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  "时间错误",
		})
	}

	allFace := modelFace.GetFaceByTime(start, end)
	if allFace == nil {
		fmt.Println("model.GetFaceByTime err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"err":  "数据库查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"data-length": len(allFace),
		"data":        allFace,
	})
}
