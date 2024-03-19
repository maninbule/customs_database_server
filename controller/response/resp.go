package response

import (
	"github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/model/modelFaceEemdding"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseBadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"err":  "请求数据或参数错误",
		"msg":  msg,
	})
}

func ResponseOKWithData(c *gin.Context, data interface{}) {
	var length int
	switch v := data.(type) {
	case []modelFace.Face:
		length = len(v)
	case []modelFaceEemdding.FaceEmbedding:
		length = len(v)
	default:
		length = 0
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"data-length": length,
		"data":        data,
	})
}
