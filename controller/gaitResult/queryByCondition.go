package gaitResult

import (
	"github.com/customs_database_server/controller/response"
	mysqlGaitResult "github.com/customs_database_server/dao/mysql/GaitResult"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"time"
)

type queryCondition struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	CameraID  string `json:"cameraID"`
}

func QueryFaceByCondition(c *gin.Context) {
	condition := queryCondition{}
	err := c.Bind(&condition)
	L1, L2 := len(condition.TimeStart), len(condition.TimeEnd)
	if err != nil || (L1+L2 > 0 && L1*L2 == 0) {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间字段不完整")
		return
	}
	query := mysqlGaitResult.CreateQuery()
	if len(condition.CameraID) > 0 {
		query = mysqlGaitResult.GetGaitByCameraId(query, condition.CameraID)
	}
	if L1+L2 > 0 && L1*L2 != 0 {
		var start, end time.Time
		ok1 := util.ParseTime(condition.TimeStart, &start)
		ok2 := util.ParseTime(condition.TimeEnd, &end)
		if !ok1 || !ok2 {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间格式不正确")
			return
		}
		query = mysqlGaitResult.GetFaceByTimeInterval(query, start, end)
	}
	if len(condition.Name) > 0 {
		query = mysqlGaitResult.GetFaceByName(query, condition.Name)
	}
	if len(condition.ID) > 0 {
		if id, err := util.ParseInt(condition.ID); err == nil {
			query = mysqlGaitResult.GetFaceById(query, id)
		} else {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "id字段需要为整数")
			return
		}
	}
	result := mysqlGaitResult.GetResult(query)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}
