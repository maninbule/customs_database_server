package faceResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlFaceResult "github.com/customs_database_server/dao/mysql/FaceResult"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	ID        string `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	TimeStart string `json:"timeStart" form:"timeStart"`
	TimeEnd   string `json:"timeEnd" form:"timeEnd"`
	CameraID  string `json:"cameraID" form:"cameraID"`
}

// QueryFaceByCondition 根据多种条件进行查询人脸识别结果
// @Summary 查询接口
// @Description 每个字段都可以为空，类似复选框
// @Tags 人脸识别结果
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param object formData queryCondition false "查询参数"
// @Success 200 {object} responseModel.Face
// @Router /face_query [post]
func QueryFaceByCondition(c *gin.Context) {
	query := getQueryFromContext(c)
	page_str := c.Param("page")
	size_str := c.Param("size")
	page, err1 := util.ParseInt(page_str)
	size, err2 := util.ParseInt(size_str)
	total := mysqlFaceResult.GetCountWithCondition(query)
	offset := (page - 1) * size
	limit := max(0, min(size, total-offset))
	if query == nil || err1 != nil || err2 != nil {
		return
	}
	result := mysqlFaceResult.GetResultWithLimit(query, offset, limit)
	for i, _ := range result {
		err1 := result[i].ConvertUTCtoLocalTime("Asia/Shanghai")
		if err1 != nil {
			response.ResponseErr(c, response.CodeErrServerErr)
		}
	}
	//result := mysqlFaceResult.GetResult(query)
	fmt.Println("GetResult = ", result)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}

func QueryFaceCountByCondition(c *gin.Context) {
	query := getQueryFromContext(c)
	if query == nil {
		return
	}
	result := mysqlFaceResult.GetCountWithCondition(query)
	fmt.Println("GetResult = ", result)
	if result == -1 {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}

func getQueryFromContext(c *gin.Context) *gorm.DB {
	err3 := c.Request.ParseForm()
	if err3 != nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return nil
	}
	condition := queryCondition{}
	fmt.Println(c.PostForm("id"))
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostForm("timeStart"))
	fmt.Println(c.PostForm("timeEnd"))
	fmt.Println(c.PostForm("cameraID"))
	err := c.ShouldBind(&condition)
	L1, L2 := len(condition.TimeStart), len(condition.TimeEnd)
	if err != nil || (L1+L2 > 0 && L1*L2 == 0) {
		fmt.Println(err, L1, L2)
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "时间字段不完整")
		return nil
	}
	fmt.Println("condition: ", condition)
	query := mysqlFaceResult.CreateQuery()
	if len(condition.CameraID) > 0 {
		query = mysqlFaceResult.GetFaceByCameraID(query, condition.CameraID)
	}
	if L1+L2 > 0 && L1*L2 != 0 {
		var start, end time.Time
		ok1 := util.ParseTime(condition.TimeStart, &start)
		ok2 := util.ParseTime(condition.TimeEnd, &end)
		fmt.Println("start = ", start)
		fmt.Println("end = ", end)
		if !ok1 || !ok2 {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间格式不正确")
			return nil
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
			return nil
		}
	}
	return query
}
