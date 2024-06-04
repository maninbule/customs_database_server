package gaitEmbedding

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/GaitEmbedding"
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
