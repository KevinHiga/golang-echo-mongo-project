package models

type Book struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Descripcion string `json:"description,omitempty" bson:"description,omitempty"`
	LibraryID   string `json:"libraryId,omitempty" bson:"libraryId,omitempty"`
	Autor       string `json:"autor,omitempty" bson:"autor,omitempty"`
	Year        string `json:"year,omitempty" bson:"year,omitempty"`
}