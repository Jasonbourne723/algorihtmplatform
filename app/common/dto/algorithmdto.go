package dto

import "time"

type AlgorithmDto struct {
	Id              int64              `json:"id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	ProjectId       int64              `json:"projectId"`
	Type            int32              `json:"type"`
	CreateUserId    int64              `json:"createUserId"`
	CreateUserName  string             `json:"createUserName"`
	CreateDate      time.Time          `json:"createDate"`
	InputAlgorithm  []AlgorithmItemDto `json:"inputAlgorithm"`
	OutputAlgorithm []AlgorithmItemDto `json:"outputAlgorithm"`
	AlgoAlgorithm   AlgorithmItemDto   `json:"algoAlgorithm"`
}

type AlgorithmItemDto struct {
	OperatorId   int64              `json:"operatorId"`
	OperatorName string             `json:"operatorName"`
	ItemParams   []KeyValuePairsDto `json:"itemParams"`
}

type KeyValuePairsDto struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddAlgorithmDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int64  `json:"projectId"`
	Type        int32  `json:"type"`
	AlgorithmItemsDto
	UserId   int64  `json:"-"`
	UserName string `json:"-"`
}

type UpdateAlgorithmDto struct {
	AddAlgorithmDto
	Id int64 `json:"id"`
}

type UpdateAlgorithmFileDto struct {
	Content string `json:"content"`
	Path    string `json:"path"`
}

type AlgorithmItemsDto struct {
	InputAlgorithm  []AlgorithmItemDto `json:"inputAlgorithm"`
	OutputAlgorithm []AlgorithmItemDto `json:"outputAlgorithm"`
	AlgoAlgorithm   AlgorithmItemDto   `json:"algoAlgorithm"`
}

type RunAlgorithm struct {
	AlgorithmId int64
	UserId      int64
}
