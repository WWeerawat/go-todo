package models

import (
	"time"
)

type User struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string      `idx:"{email},unique" json:"email" binding:"required"`
	Password  string      `json:"password" binding:"required"`
	Name      string      `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
