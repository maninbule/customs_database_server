package gaitResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlGaitResult "github.com/customs_database_server/dao/mysql/GaitResult"
	modelGaitResult "github.com/customs_database_server/model/modelGait"
	"github.com/gin-gonic/gin"
)

func CreateGaitResult(c *gin.Context) {
	fmt.Println("enter CreateGaitResult function")
	var req modelGaitResult.GaitRequest
	print(c.Request.Body)
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("GaitRequest 参数错误", err)
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}

	gait, ok := modelGaitResult.ConvertGaitRequestToGait(&req)
	if !ok {
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	ok = mysqlGaitResult.CreateGait(gait)
	if !ok {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	response.ResponseOK(c)
}
