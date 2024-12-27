package model

import "gorm.io/gorm"

type Publication struct {
	gorm.Model
	Type            string `json:"type"`
	PublicationName string `json:"publication_name"`
	Posts           []Post `json:"posts"`
}
type Post struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	PublicationId uint   `json:"publication_id"`
}