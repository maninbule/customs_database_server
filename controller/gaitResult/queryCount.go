package gaitResult

import (
	"github.com/customs_database_server/controller/response"
	mysqlGaitResult "github.com/customs_database_server/dao/mysql/GaitResult"
	"github.com/gin-gonic/gin"
)

func Getcount(c *gin.Context) {
	count := mysqlGaitResult.GaitCount()
	response.ResponseOKWithData(c, count)
}
