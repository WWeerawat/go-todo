package models

import "time"

type ToDo struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name"`
	IsDone    bool        `json:"isDone" bson:"isDone"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
	UserId    string      `json:"userId" bson:"userId"`
}
