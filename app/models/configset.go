package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type ConfigSet struct {
	Id            int64
	Name          string
	Tag           string
	Description   string
	KeyValuePairs string
	CreateDate    time.Time
	CreateUserId  int64
	IsDelete      soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
