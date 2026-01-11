package models

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if p.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return nil
}
