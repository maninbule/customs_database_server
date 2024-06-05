package model

import (
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/model/modelAttr"
	"github.com/customs_database_server/internal/model/modelFace"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
	"github.com/customs_database_server/internal/model/modelGait"
	"github.com/customs_database_server/internal/model/modelGaitEmbdding"
)

func InitModel() {
	global.DB.AutoMigrate(&modelFaceResult.Face{})
	global.DB.AutoMigrate(&modelFaceEemdding.FaceEmbedding{})
	global.DB.AutoMigrate(&modelGaitEmbdding.GaitEmbedding{})
	global.DB.AutoMigrate(&modelGaitResult.Gait{})
	global.DB.AutoMigrate(&modelAttr.Attribute{})
}
