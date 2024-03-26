package model

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelAttr"
	modelFaceResult "github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/model/modelFaceEemdding"
	modelGaitResult "github.com/customs_database_server/model/modelGait"
	"github.com/customs_database_server/model/modelGaitEmbdding"
)

func InitModel() {
	config.DB.AutoMigrate(&modelFaceResult.Face{})
	config.DB.AutoMigrate(&modelFaceEemdding.FaceEmbedding{})
	config.DB.AutoMigrate(&modelGaitEmbdding.GaitEmbedding{})
	config.DB.AutoMigrate(&modelGaitResult.Gait{})
	config.DB.AutoMigrate(&modelAttr.Attribute{})
}
