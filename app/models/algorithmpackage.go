package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type AlgorithmPackage struct {
	Id               int64
	Name             string
	AlgorithmId      int64
	AlgorithmName    string
	AlgorithmVersion string
	Config           string
	ModelFileId      int64
	CompanyId        int64
	CraneId          int64
	Status           int32
	FilePath         string
	Period           int32
	Cron             string
	CreateDate       time.Time
	IsDelete         soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
