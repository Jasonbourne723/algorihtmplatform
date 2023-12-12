package dto

import "time"

type ProjectDto struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	CreateUserId   int64     `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateDate     time.Time `json:"createDate"`
}

type AddProjectDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int64  `json:"-"`
	UserName    string `json:"-"`
}

type UpdateProjectDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
