package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, format interface{}) bool {
	b, err := c.GetRawData()
	if err != nil {
		return false
	}
	err = json.Unmarshal(b, format)
	fmt.Printf("%#v", format)
	if err != nil {
		fmt.Println("json解析失败: ", err)
		return false
	}
	return true
}
