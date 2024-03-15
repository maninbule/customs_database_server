package router

import (
	"fmt"
	"github.com/customs_database_server/controller/face"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/save_face", face.SaveFaceCompare)
	router.GET("/allFace", face.QueryAllFace)
	router.GET("/face/:startTime/:endTime", face.QueryFaceByTime)
	router.GET("/facepage/:left/:right", face.QueryFaceByLimit)
	router.POST("/face_query", face.QueryFaceByCondition)

	// 文件服务器，文件全部存储在static目录
	router.Static("/face_img/", "static")

	fmt.Println("server on port:8082.....")
	router.Run(":8082")
	return router
}
