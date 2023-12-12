package global

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/config"

	"gorm.io/gorm"
)

type Application struct {
	DB       *gorm.DB
	Config   config.Configuration
	AlgoChan chan *dto.RunAlgorithm
}

var App = new(Application)
