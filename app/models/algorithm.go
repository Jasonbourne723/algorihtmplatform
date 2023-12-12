package models

import (
	"database/sql"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Algorithm struct {
	Id             int64
	Name           string
	Remark         string
	ProjectId      int64
	Type           int32
	CreateUserId   int64
	CreateUserName string
	CreateDate     time.Time
	UpdateDate     sql.NullTime
	Version        string
	Config         string
	IsDelete       soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

type AlgorithmHistory struct {
	Id            int64
	AlgorithmId   int64
	AlgorithmName string
	ProjectId     int64
	Version       string
	Config        string
	CreateDate    time.Time
	Remark        string
}
