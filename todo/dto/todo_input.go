package dto

import "time"

type CreateToDoInput struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type UpdateToDoInput struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	UserId    string    `json:"userId,omitempty" bson:"userId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
