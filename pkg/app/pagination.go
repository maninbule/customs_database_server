// 分页逻辑

package app

import (
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/pkg/convert"
	"github.com/gin-gonic/gin"
)

// GetPage 获取路径参数中的页码
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Param("page")).MustInt()
	return max(page, 0)
}

// GetPageSize 获取路径参数中的每页item个数
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Param("size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	return min(pageSize, global.AppSetting.MaxPageSize)
}

// GetPageOffset 获取数据库的偏移量
func GetPageOffset(page, pageSize int) int {
	if page == 0 {
		return 0
	} else {
		return (page - 1) * pageSize
	}
}
