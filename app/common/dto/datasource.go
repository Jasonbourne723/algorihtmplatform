package dto

import (
	"algorithmplatform/app/common"
)

type DataSourceDto struct {
	AlgorithmId int64            `json:"algorithmId"`
	CompanyId   int64            `json:"companyId"`
	CompanyName string           `json:"companyName"`
	CraneId     int64            `json:"craneId"`
	CraneName   string           `json:"craneName"`
	BeginDate   common.TimeStamp `json:"beginDate"`
	EndDate     common.TimeStamp `json:"endDate"`
}
