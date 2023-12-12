package dto

import "time"

type OperatorDto struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	FilePath       string    `json:"filePath"`
	Description    string    `json:"description"`
	Version        string    `json:"version"`
	ClassName      string    `json:"className"`
	Type           int16     `json:"type"`
	CreateUserId   int64     `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateDate     time.Time `json:"createDate"`
	UpdateDate     time.Time `json:"updateDate"`
}

type AddOperatorDto struct {
	Name        string `json:"name"`
	FilePath    string `json:"filePath"`
	Description string `json:"description"`
	ClassName   string `json:"className"`
	Type        int16  `json:"type"`
	UserId      int64  `json:"-"`
	UserName    string `json:"-"`
}

type UpdateOperatorDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	FilePath    string `json:"filePath"`
	Description string `json:"description"`
	ClassName   string `json:"className"`
	Type        int16  `json:"type"`
}

type UpdateFileDto struct {
	OperatorId int64  `json:"operatorId"`
	Content    string `json:"content"`
	Path       string `json:"path"`
}
