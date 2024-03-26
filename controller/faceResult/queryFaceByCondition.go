package faceResult

import (
	"github.com/customs_database_server/controller/response"
	mysqlFaceResult "github.com/customs_database_server/dao/mysql/FaceResult"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"time"
)

/*
name: 名字
time_interval: 时间区间
camera_position：摄像头机位

{
	"id":"",
    "name":"",
    "timeStart":"",
    "timeEnd":"2024-3-12 00:00:00",
    "cameraID":"1"
}

按照摄像头id，时间区间，名字，id依次进行过滤查询
*/

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
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "时间字段不完整")
		return
	}
	query := mysqlFaceResult.CreateQuery()
	if len(condition.CameraID) > 0 {
		query = mysqlFaceResult.GetFaceByCameraID(query, condition.CameraID)
	}
	if L1+L2 > 0 && L1*L2 != 0 {
		var start, end time.Time
		ok1 := util.ParseTime(condition.TimeStart, &start)
		ok2 := util.ParseTime(condition.TimeEnd, &end)
		if !ok1 || !ok2 {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间格式不正确")
			return
		}
		query = mysqlFaceResult.GetFaceByTimeInterval(query, start, end)
	}
	if len(condition.Name) > 0 {
		query = mysqlFaceResult.GetFaceByName(query, condition.Name)
	}
	if len(condition.ID) > 0 {
		if id, err := util.ParseInt(condition.ID); err == nil {
			query = mysqlFaceResult.GetFaceById(query, id)
		} else {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "id字段需要为整数")
			return
		}
	}
	result := mysqlFaceResult.GetResult(query)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}
