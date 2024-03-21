package faceResult

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelFace"
	"github.com/gin-gonic/gin"
)

func GetFaceCount(c *gin.Context) {
	cnt := modelFace.GetCount()
	response.ResponseOKWithData(c, cnt)
}
