package models

import (
	"database/sql"
	"time"
)

type Operator struct {
	Id             int64
	Name           string
	FilePath       string
	Description    string
	Version        string
	ClassName      string
	Type           int16
	CreateUserId   int64
	CreateUserName string
	CreateDate     time.Time
	UpdateDate     sql.NullTime
}

type OperatorHistory struct {
	Id           int64
	OperatorId   int64
	OperatorName string
	FilePath     string
	Version      string
	CreateDate   time.Time
	Remark       string
}
