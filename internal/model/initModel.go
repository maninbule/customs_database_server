package model

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/internal/model/modelAttr"
	"github.com/customs_database_server/internal/model/modelFace"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
	"github.com/customs_database_server/internal/model/modelGait"
	"github.com/customs_database_server/internal/model/modelGaitEmbdding"
)

func InitModel() {
	config.DB.AutoMigrate(&modelFaceResult.Face{})
	config.DB.AutoMigrate(&modelFaceEemdding.FaceEmbedding{})
	config.DB.AutoMigrate(&modelGaitEmbdding.GaitEmbedding{})
	config.DB.AutoMigrate(&modelGaitResult.Gait{})
	config.DB.AutoMigrate(&modelAttr.Attribute{})
}
