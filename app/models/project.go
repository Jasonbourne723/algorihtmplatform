package models

import "time"

type Project struct {
	Id             int64
	Name           string
	Description    string
	CreateUserId   int64
	CreateUserName string
	CreateDate     time.Time
	IsDelete       int16
}
