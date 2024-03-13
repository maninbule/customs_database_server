package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

func ParseTime(timeStr string, t *time.Time) bool {
	result, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return false
	}
	t = &result
	return true
}

func ParseInt(intStr string) (int64, error) {
	result, err := strconv.ParseInt(intStr, 0, 0)
	if err != nil {
		return 0, errors.New("int转换错误")
	}
	return result, nil
}
