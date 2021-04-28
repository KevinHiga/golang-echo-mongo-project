package models

type LibraryBook struct {
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Direction        string `json:"direction,omitempty" bson:"direction,omitempty"`
	AvatarUrl        string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	Year        string `json:"year,omitempty" bson:"year,omitempty"`
	Library Book   `json:"library,omitempty" bson:"library,omitempty"`
}