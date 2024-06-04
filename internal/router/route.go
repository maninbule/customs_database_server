package router

import (
	_ "github.com/customs_database_server/docs"
	"github.com/customs_database_server/global"
	AttrResult2 "github.com/customs_database_server/internal/controller/AttrResult"
	"github.com/customs_database_server/internal/controller/demo"
	faceEmbedding2 "github.com/customs_database_server/internal/controller/faceEmbedding"
	faceResult2 "github.com/customs_database_server/internal/controller/faceResult"
	gaitEmbedding2 "github.com/customs_database_server/internal/controller/gaitEmbedding"
	gaitResult2 "github.com/customs_database_server/internal/controller/gaitResult"
	"github.com/customs_database_server/internal/controller/kafka"
	logicfaceEmbedding "github.com/customs_database_server/logic/faceEmbedding"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"net/http"
	"time"

	gs "github.com/swaggo/gin-swagger"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
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

	router.GET("/loadImageToKafka", Controllerkafka.PushImgageToKafka)
	// 人脸结果相关url
	router.POST("/save_face", faceResult2.SaveFaceCompare)
	router.GET("/allFace", faceResult2.QueryAllFace)
	router.GET("/faceImgDataBase/:startTime/:endTime", faceResult2.QueryFaceByTime)
	router.GET("/facepage/:left/:right", faceResult2.QueryFaceByLimit)
	router.POST("/face_query/:page/:size", faceResult2.QueryFaceByCondition)
	router.POST("/face_query_count", faceResult2.QueryFaceCountByCondition)
	router.GET("/face_count", faceResult2.GetFaceCount)
	router.GET("/face_result_demo", demo.ShowAllFaceResult)
	router.GET("/face_embedding_demo", demo.ShowAllFaceEmbedding)
	router.GET("/html_result_demo", demo.ShowHTMLFaceResult)

	// 人脸向量相关url
	router.GET("/init_face_embedding_database", logicfaceEmbedding.Enter)
	router.POST("/register_face", faceEmbedding2.SaveFaceEmbedding)
	router.GET("/query_all_face", faceEmbedding2.GetAllFaceEmbedding)

	// 步态向量相关url
	router.POST("/register_gait", gaitEmbedding2.SaveGaitEmbedding)
	router.GET("/get_all_gait", gaitEmbedding2.GetAllGaitEmbedding)

	// 步态识别结果url
	router.POST("/create_gait_result", gaitResult2.CreateGaitResult)
	router.POST("/query_gait_result/:page/:size", gaitResult2.QueryFaceByCondition)
	router.POST("/query_gait_count", gaitResult2.GetCountWithCondition)
	router.GET("/gait_count", gaitResult2.Getcount)

	// 高抗伪相关url
	router.POST("/create_attr_result", AttrResult2.SaveAttr)
	router.POST("/query_attr_result/:page/:size", AttrResult2.QueryFaceByCondition)
	router.POST("/query_count_with_condition", AttrResult2.QueryCountByCondition)
	// 文件服务器，文件全部存储在static目录
	router.Static("/face_img/", "static")
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	return router
}

func InitRouter() {
	router := newRouter()
	gin.SetMode(global.ServerSetting.RunMode)
	server := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic("服务器停止....")
	}
}
