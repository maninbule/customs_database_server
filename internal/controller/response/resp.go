package response

import (
	"github.com/customs_database_server/internal/model/modelAttr"
	"github.com/customs_database_server/internal/model/modelFace"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
	"github.com/customs_database_server/internal/model/modelGait"
	"github.com/customs_database_server/internal/model/modelGaitEmbdding"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeErr int64

const (
	CodeErrRequest = 1000 + iota
	CodeErrReQuestTooLarge
	CodeErrRequestParamNotExisted
	CodeErrServerErr
	CodeErrDataBase
	CodeErrDataNotExisted
	CodeErrExistedFace
)

var codeMap = map[CodeErr]string{
	CodeErrRequest:                "错误的请求",
	CodeErrReQuestTooLarge:        "Post请求数据量过大",
	CodeErrRequestParamNotExisted: "缺少请求参数",
	CodeErrServerErr:              "服务器处理错误",
	CodeErrDataBase:               "数据库错误",
	CodeErrDataNotExisted:         "请求的数据不存在",
	CodeErrExistedFace:            "重复的人脸",
}

func (c CodeErr) Msg() string {
	return codeMap[c]
}

func ResponseErr(c *gin.Context, code CodeErr) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"err":  code.Msg(),
	})
}

func ResponseErrWithMsg(c *gin.Context, code CodeErr, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"err":  code.Msg(),
		"msg":  msg,
	})
}

func ResponseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "请求成功",
	})
}

func ResponseOKWithData(c *gin.Context, data interface{}) {
	var length int
	switch v := data.(type) {
	case []modelFaceEemdding.FaceEmbedding:
		length = len(v)
	case []modelFaceResult.Face:
		length = len(v)
	case []modelGaitEmbdding.GaitEmbedding:
		length = len(v)
	case []modelGaitResult.Gait:
		length = len(v)
	case []modelAttr.Attribute:
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
