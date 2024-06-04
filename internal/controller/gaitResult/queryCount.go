package gaitResult

import (
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/GaitResult"
	"github.com/gin-gonic/gin"
)

func Getcount(c *gin.Context) {
	count := mysqlGaitResult.GaitCount()
	response.ResponseOKWithData(c, count)
}
