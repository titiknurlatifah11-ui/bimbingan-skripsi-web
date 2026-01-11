package entities

import "time"

type Products struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Category    Category
	Stock       int64     `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
