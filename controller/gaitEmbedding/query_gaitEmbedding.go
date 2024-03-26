package gaitEmbedding

import (
	"github.com/customs_database_server/controller/response"
	mysqlGaitEmbedding "github.com/customs_database_server/dao/mysql/GaitEmbedding"
	"github.com/gin-gonic/gin"
)

func GetAllGaitEmbedding(c *gin.Context) {
	allGait := mysqlGaitEmbedding.GetAllGait()
	if allGait == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	response.ResponseOKWithData(c, allGait)
}
