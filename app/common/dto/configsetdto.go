package dto

import "time"

type ConfigSetDto struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Tag           string    `json:"tag"`
	Description   string    `json:"description"`
	KeyValuePairs string    `json:"keyValuePairs"`
	CreateDate    time.Time `json:"createDate"`
}

type AddConfigSetDto struct {
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Description   string `json:"description"`
	KeyValuePairs string `json:"keyValuePairs"`
	UserId        int64  `json:"-"`
	UserName      string `json:"-"`
}

type UpdateConfigSetDto struct {
	Id int64 `json:"id"`
	AddConfigSetDto
}
