package gaitResult

import (
	"github.com/customs_database_server/controller/response"
	"github.com/customs_database_server/model/modelGait"
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
		response.ResponseBadRequest(c, "时间字段不完整")
		return
	}
	query := modelGait.CreateQuery()
	if len(condition.CameraID) > 0 {
		query = modelGait.GetGaitByCameraId(query, condition.CameraID)
	}
	if L1+L2 > 0 && L1*L2 != 0 {
		var start, end time.Time
		ok1 := util.ParseTime(condition.TimeStart, &start)
		ok2 := util.ParseTime(condition.TimeEnd, &end)
		if !ok1 || !ok2 {
			response.ResponseBadRequest(c, "时间转换错误")
			return
		}
		query = modelGait.GetFaceByTimeInterval(query, start, end)
	}
	if len(condition.Name) > 0 {
		query = modelGait.GetFaceByName(query, condition.Name)
	}
	if len(condition.ID) > 0 {
		if id, err := util.ParseInt(condition.ID); err == nil {
			query = modelGait.GetFaceById(query, id)
		} else {
			response.ResponseBadRequest(c, "id字段需要为整数")
			return
		}
	}
	result := modelGait.GetResult(query)
	if result == nil {
		response.ResponseBadRequest(c, "数据库错误")
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}
