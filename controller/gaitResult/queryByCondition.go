package gaitResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlGaitResult "github.com/customs_database_server/dao/mysql/GaitResult"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

// queryCondition 用于查询结果的请求参数
type queryCondition struct {
	ID        string `json:"id" example:"123"`                        // ID 是一个用于查询的唯一标识符
	Name      string `json:"name" example:"小明"`                       // Name 是一个用于查询的姓名
	TimeStart string `json:"timeStart" example:"2024-05-17 10:00:00"` // TimeStart 表示查询开始时间
	TimeEnd   string `json:"timeEnd" example:"2024-05-17 18:00:00"`   // TimeEnd 表示查询结束时间
	CameraID  string `json:"cameraID" example:"camera123"`            // CameraID 表示摄像头的唯一标识符
}

// QueryFaceByCondition 根据多种条件进行查询步态识别结果
// @Summary 条件查询的步态识别结果分页查询接口
// @Description 每个字段都可以为空，类似复选框
// @Tags 条件查询的步态识别结果分页查询
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param page path int true "页码"
// @Param size path int true "每页数量"
// @Param object formData queryCondition false "查询参数"
// @Success 200 {object} []modelGaitResult.Gait
// @Router /query_gait_result/:page/:size [post]
func QueryFaceByCondition(c *gin.Context) {
	query := getQuery(c)
	page_str := c.Param("page")
	size_str := c.Param("size")
	page, err1 := util.ParseInt(page_str)
	size, err2 := util.ParseInt(size_str)
	total := mysqlGaitResult.GetCountWithCondition(query)
	offset := (page - 1) * size
	limit := max(0, min(size, total-offset))
	if query == nil || err1 != nil || err2 != nil {
		return
	}
	result := mysqlGaitResult.GetResultWithLimit(query, offset, limit)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}

// GetCountWithCondition 根据多种条件进行查询步态识别结果
// @Summary 条件查询的步态识别结果个数查询接口
// @Description 每个字段都可以为空，类似复选框
// @Tags 条件查询的步态识别结果个数查询
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param object formData queryCondition false "查询参数"
// @Success 200 {object} int
// @Router /query_gait_count [post]
func GetCountWithCondition(c *gin.Context) {
	query := getQuery(c)
	if query == nil {
		return
	}
	result := mysqlGaitResult.GetCountWithCondition(query)
	fmt.Println("GetResult = ", result)
	if result == -1 {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}

func getQuery(c *gin.Context) *gorm.DB {
	condition := &queryCondition{}
	err := c.Bind(&condition)
	L1, L2 := len(condition.TimeStart), len(condition.TimeEnd)
	if err != nil || (L1+L2 > 0 && L1*L2 == 0) {
		response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间字段不完整")
		return nil
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
			return nil
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
			return nil
		}
	}
	return query
}
