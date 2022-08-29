// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// FileEntity -.
type FileEntity struct {
	Id          int16  `json:"id"       	example:"fileId"`
	Name        string `json:"name"       	example:"fileName"`
	Description string `json:"description"  example:"some text here"`
	Path        string `json:"path"     	example:"path to file"`
}
