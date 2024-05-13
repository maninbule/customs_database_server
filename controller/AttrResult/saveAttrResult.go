package AttrResult

import (
	"fmt"
	"github.com/customs_database_server/controller/response"
	mysqlAttrResult "github.com/customs_database_server/dao/mysql/AttrResult"
	"github.com/customs_database_server/model/modelAttr"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"os"
)

func paramRequest(c *gin.Context) *modelAttr.AttributeJson {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ResponseErr(c, response.CodeErrRequest)
		return nil
	}
	var attr modelAttr.AttributeJson
	if err := c.ShouldBind(&attr); err != nil {
		fmt.Println("paramRequest : shouldBind错误")
		fmt.Println("错误 = ", err)
		response.ResponseErr(c, response.CodeErrRequest)
		return nil
	}
	faceImg, err := c.FormFile("face_img")
	if err != nil {
		response.ResponseErrWithMsg(c, response.CodeErrRequestParamNotExisted, "传输的数据不存在对应face_img的文件")
		return nil
	}
	path, err := util.SaveAttrDataBase(c, faceImg)
	if err != nil {
		response.ResponseErr(c, response.CodeErrDataBase)
		return nil
	}
	attr.FaceImg = path
	return &attr
}

func SaveAttr(c *gin.Context) {
	fmt.Println("enter function SaveAttr")
	attr := paramRequest(c)
	if attr == nil {
		fmt.Println("SaveAttr : 请求转换成结构体不成功")
		//response.ResponseErr(c, response.CodeErrRequest)
		return
	}
	attribute := modelAttr.AttributeJsonToAttribute(attr)
	ok := mysqlAttrResult.CreateAttr(attribute)
	if !ok {
		fmt.Println("attribute 存储不成功")
		response.ResponseErr(c, response.CodeErrDataBase)
		err := os.Remove(attr.FaceImg)
		if err != nil {
			fmt.Println("attribute 删除文件错误")
			return
		}
	} else {
		response.ResponseOKWithData(c, gin.H{"img_url": attr.FaceImg})
	}
}
