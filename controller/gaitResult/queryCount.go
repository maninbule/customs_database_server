package gaitResult

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelGait"
	"github.com/gin-gonic/gin"
)

func Getcount(c *gin.Context) {
	count := modelGait.GaitCount()
	response.ResponseOKWithData(c, count)
}
