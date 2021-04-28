package models

type Library struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Direction string `json:"direction,omitempty" bson:"direction,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	Year      string `json:"year,omitempty" bson:"year,omitempty"`
}