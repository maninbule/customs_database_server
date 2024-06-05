package mysqlFaceEmbedding

import (
	"fmt"
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
)

func CreateFace(face *modelFaceEemdding.FaceEmbedding) bool {
	if GetCountById(int(*face.FaceId)) >= 1 {
		return false
	}
	create := global.DB.Create(face)
	if create.Error != nil {
		fmt.Println(create.Error)
		return false
	}
	return true
}

func GetAllFace() []modelFaceEemdding.FaceEmbedding {
	allFace := make([]modelFaceEemdding.FaceEmbedding, 0)
	find := global.DB.Model(&modelFaceEemdding.FaceEmbedding{}).Find(&allFace)
	if find.Error != nil {
		panic("sql执行错误[获取人脸向量数据失败]： GetallFace")
		return nil
	}
	return allFace
}

func GetCountById(id int) int {
	ans := 0
	count := global.DB.Model(&modelFaceEemdding.FaceEmbedding{}).Where("faceId = ?", id).Count(&ans)
	if count.Error != nil {
		panic("sql执行错误[获取对应ID个数是失败]： GetCountById")
		return 0
	}
	return ans
}

func GetCount() int {
	ans := 0
	count := global.DB.Model(&modelFaceEemdding.FaceEmbedding{}).Count(&ans)
	if count.Error != nil {
		panic("sql执行错误[获取对应ID个数是失败]： GetCountById")
		return 0
	}
	return ans
}
