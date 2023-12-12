package dto

type FileDto struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	IsDirectory bool      `json:"isDirectory"`
	Children    []FileDto `json:"children"`
}
