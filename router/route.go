package router

import (
	"fmt"
	"github.com/customs_database_server/controller/AttrResult"
	"github.com/customs_database_server/controller/demo"
	"github.com/customs_database_server/controller/faceEmbedding"
	"github.com/customs_database_server/controller/faceResult"
	"github.com/customs_database_server/controller/gaitEmbedding"
	"github.com/customs_database_server/controller/gaitResult"
	Controllerkafka "github.com/customs_database_server/controller/kafka"
	logicfaceEmbedding "github.com/customs_database_server/logic/faceEmbedding"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"path/filepath"
	"time"

	_ "github.com/customs_database_server/docs"

	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 配置CORS中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有域的跨域请求，出于安全考虑，应指定具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://www.example.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	templatesPath := filepath.Join("static", "templates")
	router.LoadHTMLGlob(filepath.Join(templatesPath, "*.html"))

	router.GET("/loadImageToKafka", Controllerkafka.PushImgageToKafka)
	// 人脸结果相关url
	router.POST("/save_face", faceResult.SaveFaceCompare)
	router.GET("/allFace", faceResult.QueryAllFace)
	router.GET("/faceImgDataBase/:startTime/:endTime", faceResult.QueryFaceByTime)
	router.GET("/facepage/:left/:right", faceResult.QueryFaceByLimit)
	router.POST("/face_query/:page/:size", faceResult.QueryFaceByCondition)
	router.POST("/face_query_count", faceResult.QueryFaceCountByCondition)
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
	router.POST("/query_gait_result/:page/:size", gaitResult.QueryFaceByCondition)
	router.POST("/query_gait_count", gaitResult.GetCountWithCondition)
	router.GET("/gait_count", gaitResult.Getcount)

	// 高抗伪相关url
	router.POST("/create_attr_result", AttrResult.SaveAttr)
	router.POST("/query_attr_result/:page/:size", AttrResult.QueryFaceByCondition)
	router.POST("/query_count_with_condition", AttrResult.QueryCountByCondition)
	// 文件服务器，文件全部存储在static目录
	router.Static("/face_img/", "static")
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	fmt.Println("server on port:8082.....")
	router.Run(":8082")
	return router
}
