package gaitEmbedding

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelGaitEmbdding"
	"github.com/gin-gonic/gin"
)

func GetAllGaitEmbedding(c *gin.Context) {
	allGait := modelGaitEmbdding.GetAllGait()
	if allGait == nil {
		response.ResponseInternalErr(c, "获取步态数据失败")
		return
	}
	response.ResponseOKWithData(c, allGait)
}
