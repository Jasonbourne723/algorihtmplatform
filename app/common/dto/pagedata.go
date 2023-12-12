package dto

type PageData[T any] struct {
	PageCount   int64 `json:"pageCount"`
	RecordCount int64 `json:"recordCount"`
	TModel      []T   `json:"tModel"`
}
