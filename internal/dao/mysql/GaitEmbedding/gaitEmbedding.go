package mysqlGaitEmbedding

import (
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/model/modelGaitEmbdding"
)

func CreateGait(gait *modelGaitEmbdding.GaitEmbedding) bool {
	create := global.DB.Model(&modelGaitEmbdding.GaitEmbedding{}).Create(gait)
	if create.Error != nil {
		panic("数据库创建步态信息错误")
		return false
	}
	return true
}

func GetAllGait() []modelGaitEmbdding.GaitEmbedding {
	allGait := make([]modelGaitEmbdding.GaitEmbedding, 0)
	find := global.DB.Model(&modelGaitEmbdding.GaitEmbedding{}).Find(&allGait)
	if find.Error != nil {
		panic("获取步态库失败")
		return nil
	}
	return allGait
}
