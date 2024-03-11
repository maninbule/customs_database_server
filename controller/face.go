package controller

import (
	"fmt"
	"github.com/customs_database_server/model"
	"github.com/customs_database_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateFace(c *gin.Context) {
	face := model.Face{}
	json_format := struct {
		FaceId   string `json:"face_id"`
		Name     string `json:"name"`
		FaceImg  string `json:"face_Img"`
		FaceTime string `json:"face_Time"`
	}{}
	if ok := util.ParseBody(c, &json_format); !ok {
		fmt.Println("111")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  "传入json的格式不正确或者字段为空",
		})
	}
	parseUint, err2 := strconv.ParseUint(json_format.FaceId, 0, 0)
	if err2 != nil {
		fmt.Println("Error parsing id:", err2)
		return
	}
	face.ID = uint(parseUint)
	face.Name = json_format.Name
	face.FaceImg = json_format.FaceImg
	t, err := time.Parse("2006-01-02 15:04:05", json_format.FaceTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	face.FaceTime = t
	ok := model.CreateFace(&face)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusOK,
			"err":  "入库失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "存入数据库成功",
	})
}
