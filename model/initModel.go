package model

import (
	"github.com/customs_database_server/config"
	"github.com/customs_database_server/model/modelAttr"
	"github.com/customs_database_server/model/modelFace"
	"github.com/customs_database_server/model/modelFaceEemdding"
)

func InitModel() {
	config.DB.AutoMigrate(&modelFace.Face{})
	config.DB.AutoMigrate(&modelFaceEemdding.FaceEmbedding{})
	config.DB.AutoMigrate(&modelAttr.Attribute{})
}
