package AttrResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlAttrResult "github.com/customs_database_server/dao/mysql/AttrResult"
	"github.com/customs_database_server/model/modelAttr"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"time"
)

func QueryFaceByCondition(c *gin.Context) {
	err2 := c.Request.ParseForm()
	if err2 != nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	condition := modelAttr.QueryCondition{}
	fmt.Println(c.PostForm("id"))
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostForm("mask"))
	fmt.Println(c.PostForm("hat"))
	fmt.Println(c.PostForm("glasses"))
	fmt.Println(c.PostForm("timeStart"))
	fmt.Println(c.PostForm("timeEnd"))
	fmt.Println(c.PostForm("cameraID"))
	err := c.ShouldBind(&condition)
	L1, L2 := len(condition.TimeStart), len(condition.TimeEnd)
	if err != nil || (L1+L2 > 0 && L1*L2 == 0) {
		fmt.Println(err, L1, L2)
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "时间字段不完整")
		return
	}
	fmt.Println("condition: ", condition)
	query := mysqlAttrResult.CreateQuery()
	if len(condition.CameraID) > 0 {
		query = mysqlAttrResult.GetByCameraId(query, condition.CameraID)
	}
	if L1+L2 > 0 && L1*L2 != 0 {
		var start, end time.Time
		ok1 := util.ParseTime(condition.TimeStart, &start)
		ok2 := util.ParseTime(condition.TimeEnd, &end)
		fmt.Println("start = ", start)
		fmt.Println("end = ", end)
		if !ok1 || !ok2 {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "时间格式不正确")
			return
		}
		query = mysqlAttrResult.GetByTimeInterval(query, start, end)
	}
	if condition.Mask == "yes" {
		query = mysqlAttrResult.GetByMask(query, 1)
	} else if condition.Mask == "no" {
		query = mysqlAttrResult.GetByMask(query, 0)
	}
	if condition.Hat == "yes" {
		query = mysqlAttrResult.GetByHat(query, 1)
	} else if condition.Hat == "no" {
		query = mysqlAttrResult.GetByHat(query, 0)
	}
	if condition.Glasses == "yes" {
		query = mysqlAttrResult.GetByGlasses(query, 1)
	} else if condition.Glasses == "no" {
		query = mysqlAttrResult.GetByGlasses(query, 0)
	}

	if len(condition.Name) > 0 {
		query = mysqlAttrResult.GetByName(query, condition.Name)
	}
	if len(condition.ID) > 0 {
		if id, err := util.ParseInt(condition.ID); err == nil {
			query = mysqlAttrResult.GetById(query, id)
		} else {
			response.ResponseErrWithMsg(c, response.CodeErrRequest, "id字段需要为整数")
			return
		}
	}
	result := mysqlAttrResult.GetResult(query)
	fmt.Println("GetResult = ", result)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}
