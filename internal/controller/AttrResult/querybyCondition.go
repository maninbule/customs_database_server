package AttrResult

import (
	"fmt"
	"github.com/customs_database_server/internal/controller/response"
	"github.com/customs_database_server/internal/dao/mysql/AttrResult"
	"github.com/customs_database_server/internal/dao/mysql/FaceResult"
	"github.com/customs_database_server/internal/model/modelAttr"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

// QueryFaceByCondition 根据多种条件进行查询伪装识别结果
// @Summary 伪装识别结果查询接口
// @Description 每个字段都可以为空，类似复选框
// @Tags 伪装识别结果分页查询
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param page path int true "页码"
// @Param size path int true "每页数量"
// @Param object formData modelAttr.QueryCondition false "查询参数"
// @Success 200 {object} []modelAttr.Attribute
// @Router /query_attr_result/:page/:size [post]
func QueryFaceByCondition(c *gin.Context) {
	query := getQuery(c)
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
	result := mysqlAttrResult.GetResultWithLimit(query, offset, limit)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	}
	fmt.Println("高抗伪条件查询的结果GetResult = ", result)
	if result == nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, result)
		return
	}
}

// QueryCountByCondition 根据多种条件进行查询伪装识别结果
// @Summary 伪装识别结果的总个数查询接口
// @Description 每个字段都可以为空，类似复选框
// @Tags 伪装识别结果个数查询
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param object formData modelAttr.QueryCondition false "查询参数"
// @Success 200 {object} int
// @Router /query_count_with_condition [post]
func QueryCountByCondition(c *gin.Context) {
	query := getQuery(c)
	if query == nil {
		return
	}
	count := mysqlAttrResult.GetResultCount(query)
	fmt.Println("高抗伪条件查询的结果个数为：", count)
	if count == -1 {
		response.ResponseErr(c, response.CodeErrDataBase)
		return
	} else {
		response.ResponseOKWithData(c, count)
		return
	}
}

func getQuery(c *gin.Context) *gorm.DB {
	err2 := c.Request.ParseForm()
	if err2 != nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return nil
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
		return nil
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
			return nil
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
			return nil
		}
	}
	return query
}
