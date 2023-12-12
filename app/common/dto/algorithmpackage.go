package dto

import "time"

type AlgorithmPackageDto struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	AlgorithmId   int64  `json:"algorithmId"`
	AlgorithmName string `json:"algorithmName"`
	CompanyId     int64  `json:"companyId"`
	CraneId       int64  `json:"craneId"`
	Period        int32  `json:"period"`
	Cron          string `json:"cron"`
	// CreateUserId    int64              `json:"createUserId"`
	// CreateUserName  string             `json:"createUserName"`
	CreateDate      time.Time          `json:"createDate"`
	InputAlgorithm  []AlgorithmItemDto `json:"inputAlgorithm"`
	OutputAlgorithm []AlgorithmItemDto `json:"outputAlgorithm"`
	AlgoAlgorithm   AlgorithmItemDto   `json:"algoAlgorithm"`
}

type AddAlgorithmPackageDto struct {
	Name            string             `json:"name"`
	AlgorithmId     int64              `json:"AlgorithmId"`
	AlgorithmName   string             `json:"AlgorithmName"`
	CompanyId       int64              `json:"CompanyId"`
	CraneId         int64              `json:"craneId"`
	Period          int32              `json:"period"`
	Cron            string             `json:"cron"`
	InputAlgorithm  []AlgorithmItemDto `json:"inputAlgorithm"`
	OutputAlgorithm []AlgorithmItemDto `json:"outputAlgorithm"`
	AlgoAlgorithm   AlgorithmItemDto   `json:"algoAlgorithm"`
	UserId          int64              `json:"-"`
	UserName        string             `json:"-"`
}

type AlgoithmPackageParams struct {
	Para      AlgorithmParam
	CompanyId int64
	CraneId   int64
	Perild    int32
	Cron      string
}
