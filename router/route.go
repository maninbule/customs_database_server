package router

import (
	"fmt"
	"github.com/customs_database_server/controller/faceEmbedding"
	"github.com/customs_database_server/controller/faceResult"
	"github.com/customs_database_server/controller/gaitEmbedding"
	"github.com/customs_database_server/controller/gaitResult"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 人脸结果相关url
	router.POST("/save_face", faceResult.SaveFaceCompare)
	router.GET("/allFace", faceResult.QueryAllFace)
	router.GET("/faceImgDataBase/:startTime/:endTime", faceResult.QueryFaceByTime)
	router.GET("/facepage/:left/:right", faceResult.QueryFaceByLimit)
	router.POST("/face_query", faceResult.QueryFaceByCondition)
	router.GET("/face_count", faceResult.GetFaceCount)

	// 人脸向量相关url
	router.POST("/register_face", faceEmbedding.SaveFaceEmbedding)
	router.GET("/query_all_face", faceEmbedding.GetAllFaceEmbedding)

	// 步态向量相关url
	router.POST("/register_gait", gaitEmbedding.SaveGaitEmbedding)
	router.GET("/get_all_gait", gaitEmbedding.GetAllGaitEmbedding)

	// 步态识别结果url
	router.POST("/query_gait", gaitResult.QueryFaceByCondition)
	router.GET("/gait_count", gaitResult.Getcount)

	// 文件服务器，文件全部存储在static目录
	router.Static("/face_img/", "static")

	fmt.Println("server on port:8082.....")
	router.Run(":8082")
	return router
}
