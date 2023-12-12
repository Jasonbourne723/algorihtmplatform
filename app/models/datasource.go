package models

import "time"

type DataSource struct {
	Id          int64
	AlgorithmId int64
	CompanyId   int64
	CompanyName string
	CraneId     int64
	CraneName   string
	BeginDate   time.Time
	EndDate     time.Time
}
