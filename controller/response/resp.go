package response

import (
	"github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/model/modelFaceEemdding"
	"github.com/customs_database_server/model/modelGait"
	"github.com/customs_database_server/model/modelGaitEmbdding"
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

func ResponseOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "请求成功",
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
	case []modelGaitEmbdding.GaitEmbedding:
		length = len(v)
	case []modelGait.Gait:
		length = len(v)
	default:
		length = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"data-length": length,
		"data":        data,
	})
}

func ResponseInternalErr(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"err":  "服务器内部错误",
		"msg":  msg,
	})
}
