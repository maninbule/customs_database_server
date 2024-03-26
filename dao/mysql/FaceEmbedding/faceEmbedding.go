package mysqlFaceEmbedding

import (
	"fmt"
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelFaceEemdding"
)

func CreateFace(face *modelFaceEemdding.FaceEmbedding) bool {
	create := config.DB.Create(face)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllFace() []modelFaceEemdding.FaceEmbedding {
	allFace := make([]modelFaceEemdding.FaceEmbedding, 0)
	find := config.DB.Model(&modelFaceEemdding.FaceEmbedding{}).Find(&allFace)
	if find.Error != nil {
		panic("sql执行错误[获取人脸向量数据失败]： GetallFace")
		return nil
	}
	return allFace
}
