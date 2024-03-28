package router

import (
	"fmt"
	"github.com/customs_database_server/controller/demo"
	"github.com/customs_database_server/controller/faceEmbedding"
	"github.com/customs_database_server/controller/faceResult"
	"github.com/customs_database_server/controller/gaitEmbedding"
	"github.com/customs_database_server/controller/gaitResult"
	logicfaceEmbedding "github.com/customs_database_server/logic/faceEmbedding"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	templatesPath := filepath.Join("static", "templates")
	router.LoadHTMLGlob(filepath.Join(templatesPath, "*.html"))

	// 人脸结果相关url
	router.POST("/save_face", faceResult.SaveFaceCompare)
	router.GET("/allFace", faceResult.QueryAllFace)
	router.GET("/faceImgDataBase/:startTime/:endTime", faceResult.QueryFaceByTime)
	router.GET("/facepage/:left/:right", faceResult.QueryFaceByLimit)
	router.POST("/face_query", faceResult.QueryFaceByCondition)
	router.GET("/face_count", faceResult.GetFaceCount)
	router.GET("/face_result_demo", demo.ShowAllFaceResult)
	router.GET("/face_embedding_demo", demo.ShowAllFaceEmbedding)
	router.GET("/html_result_demo", demo.ShowHTMLFaceResult)

	// 人脸向量相关url
	router.GET("/init_face_embedding_database", logicfaceEmbedding.Enter)
	router.POST("/register_face", faceEmbedding.SaveFaceEmbedding)
	router.GET("/query_all_face", faceEmbedding.GetAllFaceEmbedding)

	// 步态向量相关url
	router.POST("/register_gait", gaitEmbedding.SaveGaitEmbedding)
	router.GET("/get_all_gait", gaitEmbedding.GetAllGaitEmbedding)

	// 步态识别结果url
	router.POST("/create_gait_result", gaitResult.CreateGaitResult)
	router.POST("/query_gait_result", gaitResult.QueryFaceByCondition)
	router.GET("/gait_count", gaitResult.Getcount)

	// 文件服务器，文件全部存储在static目录
	router.Static("/face_img/", "static")

	fmt.Println("server on port:8082.....")
	router.Run(":8082")
	return router
}
