package faceResult

import (
	"github.com/customs_database_server/model/modelFace"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAllFace(c *gin.Context) {
	allFace := modelFace.GetAllFace()
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
