package models

type Product struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price float64 `json:"price"`
}