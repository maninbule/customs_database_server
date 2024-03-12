package face

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
	jsonFormat := struct {
		FaceId   string `json:"face_id"`
		Name     string `json:"name"`
		FaceTime string `json:"face_Time"`
		FaceImg  string `json:"face_Img"`
	}{}
	if ok := util.ParseBody(c, &jsonFormat); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  "传入json的格式不正确或者字段为空",
		})
	}
	fmt.Printf("接受到的%#v", jsonFormat)
	parseUint, err2 := strconv.ParseUint(jsonFormat.FaceId, 0, 0)
	if err2 != nil {
		fmt.Println("Error parsing id:", err2)
		return
	}
	face.FaceId = uint(parseUint)
	face.Name = jsonFormat.Name
	face.FaceImg = jsonFormat.FaceImg
	t, err := time.Parse("2006-01-02 15:04:05", jsonFormat.FaceTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	face.FaceTime = t.Add(-8 * time.Hour)
	fmt.Println("face格式化之后： ", face)
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
