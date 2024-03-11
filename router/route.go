package router

import (
	"fmt"
	"github.com/customs_database_server/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/save_face", controller.CreateFace)

	fmt.Println("server on port:8082.....")

	router.Run(":8082")
	return router
}
