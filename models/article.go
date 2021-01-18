package models

import "time"

type Article struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
